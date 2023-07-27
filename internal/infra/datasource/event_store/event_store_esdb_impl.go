package eventstore

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	eventstore "github.com/program-world-labs/DDDGo/pkg/event_store"
)

const (
	ReadStreadNum = 100
	EventSplitNum = 2
)

var _ event.Store = (*DBImpl)(nil)

type DBImpl struct {
	esdb   *eventstore.StoreDB
	mapper *event.TypeMapper
}

func NewEventStoreDBImpl(esdb *eventstore.StoreDB, mapper *event.TypeMapper) (*DBImpl, error) {
	return &DBImpl{esdb: esdb, mapper: mapper}, nil
}

func (e *DBImpl) Store(ctx context.Context, events []event.DomainEvent, version int) error {
	return e.store(ctx, events, version, false)
}

func (e *DBImpl) SafeStore(ctx context.Context, events []event.DomainEvent, expectedVersion int) error {
	return e.store(ctx, events, expectedVersion, true)
}

func (e *DBImpl) Load(ctx context.Context, streamID string, version int) ([]event.DomainEvent, error) {
	return e.loadFrom(ctx, streamID, uint64(version))
}

func (e *DBImpl) Close() error {
	return e.esdb.Close()
}

func (e *DBImpl) store(ctx context.Context, events []event.DomainEvent, version int, safe bool) error {
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

func (e *DBImpl) loadFrom(ctx context.Context, id string, version uint64) ([]event.DomainEvent, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-eventstore-loadFrom")

	defer span.End()

	streamID := id // use the provided id to construct the stream ID
	events := make([]event.DomainEvent, 0)
	// region read-from-stream-position
	ropts := esdb.ReadStreamOptions{
		From: esdb.Revision(version),
	}

	stream, err := e.esdb.Client.ReadStream(context.Background(), streamID, ropts, ReadStreadNum)
	if err != nil {
		return nil, err
	}

	defer stream.Close()

	for {
		stream, err := stream.Recv()
		if err, ok := esdb.FromError(err); !ok {
			if err.Code() == esdb.ErrorCodeResourceNotFound {
				return nil, domainerrors.WrapWithSpan(ErrorCodeResourceNotFound, err, span)
			} else if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, err
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
		if len(parts) != EventSplitNum {
			return nil, domainerrors.WrapWithSpan(ErrorCodeEventFormatWrong, err, span)
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
	}

	return events, nil
}
