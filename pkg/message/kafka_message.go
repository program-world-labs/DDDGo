package message

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

var _ event.Producer = (*KafkaMessage)(nil)

type KafkaMessage struct {
	Subscriber *kafka.Subscriber
	Publisher  *kafka.Publisher
	Tracer     *KafkaTracer
	Router     *message.Router
}

func NewKafkaMessage(brokers []string, groupID string) (*KafkaMessage, error) {
	// Create a new Kafka subscriber config.
	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               brokers,
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: kafka.DefaultSaramaSubscriberConfig(),
			ConsumerGroup:         groupID,
		},
		watermill.NewStdLogger(false, false),
	)

	if err != nil {
		return nil, err
	}

	// Create a new Kafka publisher config.
	config := kafka.DefaultSaramaSyncPublisherConfig()
	// 設置acks
	config.Producer.RequiredAcks = sarama.WaitForAll // 等於 acks=-1
	// 設置retries
	config.Producer.Retry.Max = 10 // 重試10次

	// 建立一個新的 OTLP Tracer
	tracer := NewKafkaTracer(domainerrors.GruopID)

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:               brokers,
			Marshaler:             kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: config,
			// OTELEnabled:           true,
			// Tracer:                tracer,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	// Create a new router.
	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, err
	}

	router.AddMiddleware(tracer.Trace())

	return &KafkaMessage{Subscriber: subscriber, Publisher: publisher, Tracer: tracer, Router: router}, nil
}

func (k *KafkaMessage) PublishEvent(ctx context.Context, topic string, event interface{}) error {
	// Transform event to message
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Failed to marshal event: %v", err)

		return err
	}

	// Create a new message with event bytes.
	msg := message.NewMessage(watermill.NewUUID(), eventBytes)

	// Inject context to metadata
	propagator := propagation.TraceContext{}
	header := http.Header{}
	propagator.Inject(ctx, propagation.HeaderCarrier(header))

	metadata := make(message.Metadata)

	for k, v := range header {
		if len(v) > 0 {
			// If there are multiple values, just take the first one.
			// You might want to handle this differently depending on your use case.
			metadata[k] = v[0]
		}
	}

	msg.Metadata = metadata

	// Start tracing
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, k.structName(k.Publisher))

	defer span.End()

	// 設置 metadata Event Type
	// msg.Metadata.Set("event_type", topic)

	err = k.Publisher.Publish(topic, msg)
	if err != nil {
		return err
	}

	return nil
}

func (k *KafkaMessage) Close() error {
	k.Subscriber.Close()
	k.Publisher.Close()

	return nil
}

func (k *KafkaMessage) structName(v interface{}) string {
	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}

	s := fmt.Sprintf("%T", v)
	// trim the pointer marker, if any
	return strings.TrimLeft(s, "*")
}
