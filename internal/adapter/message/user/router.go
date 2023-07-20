package user

import (
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

type Routes struct {
	e event.TypeMapper
	s user.IService
	l pwlogger.Interface
}

func NewUserRoutes(e event.TypeMapper, u user.IService, l pwlogger.Interface) *Routes {
	// Register event
	e.Register((*event.UserCreatedEvent)(nil))
	e.Register((*event.UserPasswordChangedEvent)(nil))
	e.Register((*event.UserEmailChangedEvent)(nil))

	return &Routes{e: e, s: u, l: l}
}

func (u *Routes) Handler(msg *message.Message) error {
	log.Println("userRoutes received message", msg.UUID)

	// msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
	return nil
}
