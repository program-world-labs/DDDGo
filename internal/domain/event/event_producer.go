package event

type EventProducer interface {
	// PublishEvent publish event to pubsub server.
	PublishEvent(topic string, event interface{}) error

	// Close connection.
	Close() error
}
