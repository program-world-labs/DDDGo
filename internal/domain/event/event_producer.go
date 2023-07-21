package event

import "context"

type Producer interface {
	// PublishEvent publish event to pubsub server.
	PublishEvent(ctx context.Context, topic string, event interface{}) error

	// Close connection.
	Close() error
}
