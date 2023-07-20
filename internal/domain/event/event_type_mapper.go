package event

import (
	"fmt"
	"reflect"
)

type TypeMapper struct {
	mapper map[string]reflect.Type
}

func NewEventTypeMapper() *TypeMapper {
	return &TypeMapper{
		mapper: make(map[string]reflect.Type),
	}
}

func (m *TypeMapper) Register(event interface{}) {
	t := reflect.TypeOf(event).Elem()
	m.mapper[t.Name()] = t
}

func (m *TypeMapper) NewInstance(eventName string) (event interface{}, err error) {
	if t, ok := m.mapper[eventName]; ok {
		event = reflect.New(t).Interface()
	} else {
		err = fmt.Errorf("%w: %s", ErrNewEventTypeMapper, event)
	}

	return event, err
}
