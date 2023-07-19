package event

import "context"

type EventStore interface {
	Store(ctx context.Context, events []DomainEvent, version int) error
	Load(ctx context.Context, aggregateID string, version int) ([]DomainEvent, error)
	SafeStore(ctx context.Context, events []DomainEvent, expectedVersion int) error
}
