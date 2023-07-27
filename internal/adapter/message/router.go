package message

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel/attribute"

	"github.com/program-world-labs/DDDGo/internal/adapter/message/role"
	"github.com/program-world-labs/DDDGo/internal/adapter/message/user"
	"github.com/program-world-labs/DDDGo/internal/application"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	pkg_message "github.com/program-world-labs/DDDGo/pkg/message"
)

// config represents the configuration options available for subscriber
// middlewares and publisher decorators.
type config struct {
	spanAttributes []attribute.KeyValue
}

// Option provides a convenience wrapper for simple options that can be
// represented as functions.
type Option func(*config)

// WithSpanAttributes includes the given attributes to the generated Spans.
func WithSpanAttributes(attributes ...attribute.KeyValue) Option {
	return func(c *config) {
		c.spanAttributes = attributes
	}
}

type Router struct {
	Handler     *pkg_message.KafkaMessage
	EventMapper *event.TypeMapper
	S           application.Services
	L           pwlogger.Interface
}

func NewRouter(handler *pkg_message.KafkaMessage, mapper *event.TypeMapper, s application.Services, l pwlogger.Interface) (*message.Router, error) {
	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	handler.Router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	handler.Router.AddMiddleware(
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

	// User routes
	User := user.NewUserRoutes(*mapper, s.User, l)
	Role := role.NewRoleRoutes(*mapper, s.Role, l)

	// Add event handlers
	handler.Router.AddNoPublisherHandler("User.Handler", "User", handler.Subscriber, User.Handler)
	handler.Router.AddNoPublisherHandler("Role.Handler", "Role", handler.Subscriber, Role.Handler)

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	ctx := context.Background()
	if err := handler.Router.Run(ctx); err != nil {
		panic(err)
	}

	return handler.Router, nil
}
