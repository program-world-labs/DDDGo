package event

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Event interface {
	GetEventType() string
	GetData() interface{}
	CreatedAt() time.Time

	GetAggregateType() string
	GetAggregateID() string
	GetVersion() int

	Serialize() (string, error)
	Deserialize(string) (interface{}, error)
}

type DomainEvent struct {
	ID        uuid.UUID		// 事件唯一ID
	EventType string		// 事件的Type名稱，例如：AccountCreatedEvent
	Data      interface{}	// 事件資料，例如&AccountCreatedEvent{}
	CreatedAt time.Time

	AggregateType string	// 事件所屬的Aggregate Type，例如AccountCreatedEvent屬於Account
	AggregateID   string	// aggregate儲存在資料庫，相對應的ID
	Version       int		// aggregate的版本
}

func NewDomainEvent(aggregateID, aggregateType string, version int, data interface{}) *DomainEvent {
	_, eventType := GetTypeName(data)

	uid, err := uuid.NewV4()
	if err != nil {
		return nil
	}

	e := &DomainEvent{
		ID:            uid,
		EventType:     eventType,
		Data:          data,
		CreatedAt:     time.Now(),
		AggregateType: aggregateType,
		AggregateID:   aggregateID,
		Version:       version,
	}

	return e
}

func (e *DomainEvent) GetID() uuid.UUID {
	return e.ID
}

// EventType implements the EventType method of the Event interface.
func (e *DomainEvent) GetEventType() string {
	return e.EventType
}
func (e *DomainEvent) SetEventType(et string) {
	e.EventType = et
}

// Data implements the Data method of the Event interface.
func (e *DomainEvent) GetData() interface{} {
	return e.Data
}

// Timestamp implements the Timestamp method of the Event interface.
func (e *DomainEvent) SetTimestamp() time.Time {
	return e.CreatedAt
}

// AggregateType implements the AggregateType method of the Event interface.
func (e *DomainEvent) GetAggregateType() string {
	return e.AggregateType
}

// AggregateID implements the AggregateID method of the Event interface.
func (e *DomainEvent) GetAggregateID() string {
	return e.AggregateID
}

// Version implements the Version method of the Event interface.
func (e *DomainEvent) GetVersion() int {
	return e.Version
}

func (e *DomainEvent) SetVersion(v int) {
	e.Version = v
}

// Metadata implements the Metadata method of the Event interface.
// func (e DomainEvent) Metadata() map[string]interface{} {
// 	return e.metadata
// }

// String implements the String method of the Event interface.
func (e *DomainEvent) String() string {
	str := string(e.EventType)

	if e.AggregateID != "" && e.Version != 0 {
		str += fmt.Sprintf("(%s, v%d)", e.AggregateID, e.Version)
	}

	return str
}

// GetTypeName of given struct
func GetTypeName(source interface{}) (reflect.Type, string) {
	rawType := reflect.TypeOf(source)

	// source is a pointer, convert to its value
	if rawType.Kind() == reflect.Ptr {
		rawType = rawType.Elem()
	}

	name := rawType.String()
	// we need to extract only the name without the package
	// name currently follows the format `package.StructName`
	parts := strings.Split(name, ".")
	return rawType, parts[1]
}
