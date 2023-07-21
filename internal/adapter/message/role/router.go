package role

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/application/role"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

type Routes struct {
	e event.TypeMapper
	s role.IService
	l pwlogger.Interface
}

func NewRoleRoutes(e event.TypeMapper, r role.IService, l pwlogger.Interface) *Routes {
	// Register event
	e.Register((*event.RoleCreatedEvent)(nil))
	e.Register((*event.RoleDescriptionChangedEvent)(nil))
	e.Register((*event.RolePermissionUpdatedEvent)(nil))

	return &Routes{e: e, s: r, l: l}
}

func (u *Routes) Handler(msg *message.Message) error {
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
		err = u.create(msg.Context(), eventType.(*event.RoleCreatedEvent))
		if err != nil {
			return err
		}
	default:
		return domainerrors.Wrap(ErrorCodeHandleMessage, err)
	}

	return nil
}

func (u *Routes) create(ctx context.Context, event *event.RoleCreatedEvent) error {
	// Transform event data to service input
	// info := role.CreatedInput{}
	info := role.UpdatedInput{}
	if err := copier.Copy(&info, event); err != nil {
		return domainerrors.Wrap(ErrorCodeRoleCopyToInput, err)
	}

	info.Permissions = strings.Join(event.Permissions, ",")

	// _, err := u.s.CreateRole(ctx, &info)
	_, err := u.s.UpdateRole(ctx, &info)
	if err != nil {
		return err
	}

	return nil
}
