package role

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/DDDGo/internal/application/role"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/pwlogger"
)

type RoleRoutes struct {
	e event.EventTypeMapper
	s role.IService
	l pwlogger.Interface
}

func NewRoleRoutes(e event.EventTypeMapper, r role.IService, l pwlogger.Interface) *RoleRoutes {
	// Register event
	e.Register((*event.RoleCreatedEvent)(nil))
	e.Register((*event.RoleDescriptionChangedEvent)(nil))
	e.Register((*event.RolePermissionUpdatedEvent)(nil))

	return &RoleRoutes{e: e, s: r, l: l}
}

func (u *RoleRoutes) Handler(msg *message.Message) error {
	log.Println("RoleRoutes received message", msg.UUID)

	// Transform message to domain event
	domainEvent := &event.DomainEvent{}
	err := json.Unmarshal(msg.Payload, domainEvent)
	if err != nil {
		return err
	}

	// Json Unmarshal: To Domain Event Type
	eventType, err := u.e.NewInstance(domainEvent.EventType)
	if err != nil {
		return err
	}
	// data map to JSON
	jsonData, err := json.Marshal(domainEvent.Data)
	if err != nil {
		panic(err)
	}
	// Json data to domain event
	err = json.Unmarshal(jsonData, &eventType)
	if err != nil {
		panic(err)
	}

	switch domainEvent.EventType {
	case "RoleCreatedEvent":
		u.create(eventType.(*event.RoleCreatedEvent))
	}

	// // Transform message to domain event
	// domainEvent, err := u.toDomainEvent(msg)
	// if err != nil {
	// 	return err
	// }
	// // Transform event data to service input
	// info := &role.CreatedInput{}
	// if err := copier.Copy(info, domainEvent); err != nil {
	// 	return domainerrors.Wrap(100000, err)
	// }

	// e := &event.DomainEvent{}
	// err := json.Unmarshal(msg.Payload, e)
	// if err != nil {
	// 	return err
	// }

	// // Call service
	// switch e.EventType {
	// case "RoleCreatedEvent":
	// 	// 将 map 转换为 JSON
	// 	data, err := json.Marshal(e.Data)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// 将 JSON 转换为 RoleCreatedEvent
	// 	var eventData event.RoleCreatedEvent
	// 	err = json.Unmarshal(data, &eventData)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	// Transform event data to service input
	// 	info := role.CreatedInput{}
	// 	if err := copier.Copy(info, eventData); err != nil {
	// 		return domainerrors.Wrap(100000, err)
	// 	}

	// 	_, err = u.s.CreateRole(context.Background(), &info)
	// }

	return nil
}

func (u *RoleRoutes) create(event *event.RoleCreatedEvent) error {
	// Transform event data to service input
	info := role.CreatedInput{}
	if err := copier.Copy(&info, event); err != nil {
		return domainerrors.Wrap(100000, err)
	}
	info.Permissions = strings.Join(event.Permissions, ",")

	_, err := u.s.CreateRole(context.Background(), &info)
	if err != nil {
		return err
	}

	return nil
}

func (u *RoleRoutes) toDomainEvent(msg *message.Message) (interface{}, error) {
	// Transform message to domain event
	domainEvent := &event.DomainEvent{}
	err := json.Unmarshal(msg.Payload, domainEvent)
	if err != nil {
		return nil, err
	}

	// Json Unmarshal: To Domain Event Type
	eventType, err := u.e.NewInstance(domainEvent.EventType)
	if err != nil {
		return nil, err
	}
	// data map to JSON
	jsonData, err := json.Marshal(domainEvent.Data)
	if err != nil {
		panic(err)
	}
	// Json data to domain event
	err = json.Unmarshal(jsonData, &eventType)
	if err != nil {
		panic(err)
	}

	return eventType, nil
}
