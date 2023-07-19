package eventstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"

	"github.com/program-world-labs/DDDGo/internal/domain/event"
	eventstore "github.com/program-world-labs/DDDGo/pkg/event_store"
)

var _ event.EventStore = (*EventStoreDBImpl)(nil)

type EventStoreDBImpl struct {
	esdb   *eventstore.EventStoreDB
	mapper *event.EventTypeMapper
}

func NewEventStoreDBImpl(esdb *eventstore.EventStoreDB, mapper *event.EventTypeMapper) (*EventStoreDBImpl, error) {

	return &EventStoreDBImpl{esdb: esdb, mapper: mapper}, nil
}

func (e *EventStoreDBImpl) Store(ctx context.Context, events []event.DomainEvent, version int) error {
	return e.store(ctx, events, version, false)
}

func (e *EventStoreDBImpl) SafeStore(ctx context.Context, events []event.DomainEvent, expectedVersion int) error {
	return e.store(ctx, events, expectedVersion, true)
}

func (e *EventStoreDBImpl) Load(ctx context.Context, streamID string, version int) ([]event.DomainEvent, error) {
	return e.loadFrom(ctx, streamID, uint64(version))
}

func (e *EventStoreDBImpl) Close() error {
	return e.esdb.Close()
}

func (e *EventStoreDBImpl) store(ctx context.Context, events []event.DomainEvent, version int, safe bool) error {
	if len(events) == 0 {
		return nil
	}

	// Build all event records, with incrementing versions starting from the
	// original aggregate version.
	eventsDatas := make([]esdb.EventData, len(events))
	streamID := events[0].AggregateType + "-" + events[0].AggregateID

	for i, event := range events {
		// Create the event record with timestamp.
		eventsDatas[i] = esdb.EventData{
			EventID:     event.GetID(),
			EventType:   event.GetEventType(),
			ContentType: esdb.ContentTypeJson,
		}

		// Marshal event data
		if event.Data != nil {
			rawData, err := json.MarshalIndent(event.Data, "", "\t")
			if err != nil {
				return err
			}

			eventsDatas[i].Data = rawData
		}
	}

	// Insert a new aggregate or append to an existing.
	var option esdb.AppendToStreamOptions
	if safe {
		// Append to an Exist and Validate Version
		// if version != events[len(eventsDatas)-1].GetVersion() {
		// 	return errors.New("Version not Match")
		// }
		option = esdb.AppendToStreamOptions{
			ExpectedRevision: esdb.Revision(uint64(version)),
		}
	} else {
		// Insert a New Aggregate
		option = esdb.AppendToStreamOptions{
			ExpectedRevision: esdb.Any{},
		}
	}

	_, err := e.esdb.Client.AppendToStream(ctx, streamID, option, eventsDatas...)

	return err
}

func (e *EventStoreDBImpl) loadFrom(ctx context.Context, id string, version uint64) ([]event.DomainEvent, error) {
	streamID := id // use the provided id to construct the stream ID
	events := make([]event.DomainEvent, 0)
	// region read-from-stream-position
	ropts := esdb.ReadStreamOptions{
		From: esdb.Revision(version),
	}

	stream, err := e.esdb.Client.ReadStream(context.Background(), streamID, ropts, 100)
	defer stream.Close()

	if err != nil {
		return nil, err
	}

	for {
		stream, err := stream.Recv()
		if err, ok := esdb.FromError(err); !ok {
			if err.Code() == esdb.ErrorCodeResourceNotFound {
				fmt.Print("Stream not found")
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				panic(err)
			}
		}

		// Json Unmarshal: To Domain Event Type
		eventType, err := e.mapper.NewInstance(stream.Event.EventType)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(stream.Event.Data, &eventType)
		if err != nil {
			return nil, err
		}

		// Get Aggregate Type and ID
		parts := strings.Split(stream.Event.StreamID, "-")
		if len(parts) != 2 {
			return nil, errors.New("Stream Event Invalid Format")
		}
		aggregateType := parts[0]
		aggregateID := parts[1]

		// Create Domain Event
		e := event.NewDomainEvent(
			aggregateID,
			aggregateType,
			int(stream.Event.EventNumber),
			eventType,
		)
		// Append to Events
		events = append(
			events,
			*e,
		)

		fmt.Printf("Event> %v", stream)
	}

	return events, nil

}
