package event

import (
	"fmt"
	"reflect"
)

type EventTypeMapper struct {
	mapper map[string]reflect.Type
}

func NewEventTypeMapper() *EventTypeMapper {
	return &EventTypeMapper{
		mapper: make(map[string]reflect.Type),
	}
}

func (m *EventTypeMapper) Register(event interface{}) {
	t := reflect.TypeOf(event).Elem()
	m.mapper[t.Name()] = t
}

func (m *EventTypeMapper) NewInstance(eventName string) (event interface{}, err error) {
	if t, ok := m.mapper[eventName]; ok {
		event = reflect.New(t).Interface()
	} else {
		err = fmt.Errorf("unknown event: %s", eventName)
	}
	
	return event, err
}
