package event

import "context"

type MessageHandlerFunc func(key string, message string) error

type Consumer interface {
	SubscribeEvent(ctx context.Context, topic string, handler MessageHandlerFunc) error
	Close() error
}
