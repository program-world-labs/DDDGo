package message

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

var _ event.EventProducer = (*KafkaMessage)(nil)

type KafkaMessage struct {
	Subscriber *kafka.Subscriber
	Publisher  *kafka.Publisher
	// Router     *message.Router
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
	// saramaConfig := sarama.NewConfig()
	// saramaConfig.Producer.RequiredAcks = sarama.WaitForAll // Equivalent to 'acks=all'
	// saramaConfig.Producer.Retry.Max = 2                    // Set retries to 2

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   brokers,
			Marshaler: kafka.DefaultMarshaler{},
			// OverwriteSaramaConfig: saramaConfig,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	return &KafkaMessage{Subscriber: subscriber, Publisher: publisher}, nil
}

func (k *KafkaMessage) PublishEvent(topic string, event interface{}) error {
	// Transform event to message
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Fatalf("Failed to marshal event: %v", err)

		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), eventBytes)

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
