package message

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/adapter/message/role"
	"github.com/program-world-labs/DDDGo/internal/adapter/message/user"
	"github.com/program-world-labs/DDDGo/internal/application"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	pkg_message "github.com/program-world-labs/DDDGo/pkg/message"
)

type Router struct {
	Handler     *pkg_message.KafkaMessage
	EventMapper *event.TypeMapper
	S           application.Services
	L           pwlogger.Interface
}

func NewRouter(handler *pkg_message.KafkaMessage, mapper *event.TypeMapper, s application.Services, l pwlogger.Interface) (*message.Router, error) {
	// Create a new router.
	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, err
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          watermill.NewStdLogger(false, false),
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	// ctx := context.Background()
	// if err := router.Run(ctx); err != nil {
	// 	panic(err)
	// }

	// User routes
	User := user.NewUserRoutes(*mapper, s.User, l)
	Role := role.NewRoleRoutes(*mapper, s.Role, l)

	// Add event handlers
	router.AddNoPublisherHandler("User", "User", handler.Subscriber, User.Handler)
	router.AddNoPublisherHandler("Role", "Role", handler.Subscriber, Role.Handler)

	// 创建一个通道用于接收错误
	errChan := make(chan error)

	// 启动 Router
	go func() {
		errChan <- router.Run(context.Background())
	}()

	// 通过通道接收错误
	err = <-errChan
	if err != nil {
		// 处理错误
		l.Panic().Err(err).Msg("router run error")
	}
	// 等待直到所有處理器都已啟動
	<-router.Running()

	return router, nil
}
