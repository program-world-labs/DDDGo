package event

type Producer interface {
	// PublishEvent publish event to pubsub server.
	PublishEvent(topic string, event interface{}) error

	// Close connection.
	Close() error
}
