package message

import (
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

type KafkaTracer struct {
	tracer trace.Tracer
}

func NewKafkaTracer(groupID string) *KafkaTracer {
	return &KafkaTracer{
		tracer: otel.Tracer(groupID),
	}
}

// Trace defines a middleware that will add tracing.
func (t *KafkaTracer) Trace(options ...Option) message.HandlerMiddleware {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return t.TraceHandler(h, options...)
	}
}

// TraceHandler decorates a watermill HandlerFunc to add tracing when a message is received.
func (t *KafkaTracer) TraceHandler(h message.HandlerFunc, options ...Option) message.HandlerFunc {
	tracer := otel.Tracer(domainerrors.GruopID)
	config := &config{}

	for _, opt := range options {
		opt(config)
	}

	spanOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(config.spanAttributes...),
	}

	return func(msg *message.Message) ([]*message.Message, error) {
		// Convert message.Metadata to http.Header
		header := make(http.Header)
		for k, v := range msg.Metadata {
			header.Set(k, v)
		}

		// Extract SpanContext from the header
		propagator := propagation.TraceContext{}
		ctx := propagator.Extract(msg.Context(), propagation.HeaderCarrier(header))

		spanName := message.HandlerNameFromCtx(ctx)
		ctx, span := tracer.Start(ctx, spanName, spanOptions...)
		span.SetAttributes(
			semconv.MessagingDestinationKindTopic,
			semconv.MessagingDestinationKey.String(message.SubscribeTopicFromCtx(ctx)),
			semconv.MessagingOperationReceive,
		)
		msg.SetContext(ctx)

		events, err := h(msg)

		if err != nil {
			span.RecordError(err)
		}

		span.End()

		return events, err
	}
}

// TraceNoPublishHandler decorates a watermill NoPublishHandlerFunc to add tracing when a message is received.
func (t *KafkaTracer) TraceNoPublishHandler(h message.NoPublishHandlerFunc, options ...Option) message.NoPublishHandlerFunc {
	decoratedHandler := t.TraceHandler(func(msg *message.Message) ([]*message.Message, error) {
		return nil, h(msg)
	}, options...)

	return func(msg *message.Message) error {
		_, err := decoratedHandler(msg)

		return err
	}
}
