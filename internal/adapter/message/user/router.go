package user

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/pwlogger"
)

type UserRoutes struct {
	e event.EventTypeMapper
	s user.IService
	l pwlogger.Interface
}

func NewUserRoutes(e event.EventTypeMapper, u user.IService, l pwlogger.Interface) *UserRoutes {
	// Register event
	e.Register((*event.UserCreatedEvent)(nil))
	e.Register((*event.UserPasswordChangedEvent)(nil))
	e.Register((*event.UserEmailChangedEvent)(nil))

	return &UserRoutes{s: u, l: l}
}

func (u *UserRoutes) Handler(msg *message.Message) error {
	log.Println("userRoutes received message", msg.UUID)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return nil
}
