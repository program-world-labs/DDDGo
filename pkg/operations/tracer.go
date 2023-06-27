package operations

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var _ ITracer = (*Tracer)(nil)

type ITracer interface {
	WithSpan(ctx context.Context, name string, f func(context.Context))
}

type Tracer struct {
	tracer trace.Tracer
}

func NewTracer(applicationName string) *Tracer {
	tracer := otel.Tracer(applicationName)

	return &Tracer{
		tracer: tracer,
	}
}

func (t *Tracer) WithSpan(ctx context.Context, name string, f func(context.Context)) {
	ctx, span := t.tracer.Start(ctx, name)
	defer span.End()
	f(ctx)
}
