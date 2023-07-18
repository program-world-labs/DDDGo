package aggregate

import (
	"reflect"
	"strings"

	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

type AggregateHandler interface {
	LoadFromHistory(events []event.DomainEvent) error
	ApplyEvent(event *event.DomainEvent) error
	ApplyEventHelper(aggregate AggregateHandler, event *event.DomainEvent, commit bool)
	// HandleCommand(command interface{}) error
	UnCommitedEvents() []event.DomainEvent
	ClearUnCommitedEvents()
	IncrementVersion()
	GetID() string
	GetTypeName(source interface{}) (reflect.Type, string)
	GetVersion() int
}

type BaseAggregate struct {
	// The aggregate ID
	ID string
	// The aggregate type
	Type string
	// The aggregate version
	Version int
	// The aggregate events
	Events []event.DomainEvent
}

func (b *BaseAggregate) UnCommitedEvents() []event.DomainEvent {
	return b.Events
}

func (b *BaseAggregate) ClearUnCommitedEvents() {
	b.Events = []event.DomainEvent{}
}

func (b *BaseAggregate) IncrementVersion() {
	b.Version++
}

func (b *BaseAggregate) GetVersion() int {
	return b.Version
}

func (b *BaseAggregate) GetID() string {
	return b.ID
}

func (b *BaseAggregate) ApplyEventHelper(aggregate AggregateHandler, event *event.DomainEvent, commit bool) {
	// increments the version in event and aggregate
	b.IncrementVersion()

	// set the aggregate type
	_, aggregateType := b.GetTypeName(aggregate)
	b.Type = aggregateType

	// apply the event itself
	aggregate.ApplyEvent(event)

	// Check if Need to commit to EventStore and EventPublisher
	if commit {
		// add the event to the list of uncommitted events
		event.SetVersion(b.Version)
		_, et := b.GetTypeName(event.Data)
		event.SetEventType(et)
		b.Events = append(b.Events, *event)
	}
}

func (b *BaseAggregate) GetTypeName(source interface{}) (reflect.Type, string) {
	rawType := reflect.TypeOf(source)

	// source is a pointer, convert to its value
	if rawType.Kind() == reflect.Ptr {
		rawType = rawType.Elem()
	}

	name := rawType.String()
	// we only need the name, not the package
	// the name follows the format `package.StructName`
	parts := strings.Split(name, ".")
	
	return rawType, parts[len(parts)-1]
}
