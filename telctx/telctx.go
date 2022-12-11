package telctx

import (
	"context"

	"github.com/brendoncarroll/stdctx/units"
)

type Incrementable interface {
	int | int8 | int16 | int32 | int64
	uint | uint8 | uint16 | uint32 | uint64
	float32 | float64
}

// Incr increments the metric m by x.
func Incr[T Incrementable](ctx context.Context, m string, x T, u units.Unit) {
	c := FromContext(ctx)
	c.Incr(m, x, u)
}

type Markable interface {
	Incrementable
	string
}

// Mark sets the metric m to equal x.
func Mark[T Markable](ctx context.Context, m string, x T, u units.Unit) {
	c := FromContext(ctx)
	c.Set(m, x, u)
}

// Collector is used to collect metrics
type Collector interface {
	Incr(m string, x any, u units.Unit)
	Set(m string, x any, u units.Unit)
	// SetDenom(m string, x any, u units.Unit)
}

type NullCollector struct{}

func (c NullCollector) Incr(m string, x any, u units.Unit) {}

func (c NullCollector) Set(m string, x any, u units.Unit) {}

// func (c NullCollector) SetDenom(m string, x any, u units.Unit) {}

type prefixedCollector struct {
	inner  Collector
	prefix string
}

func (c prefixedCollector) Incr(m string, x any, u units.Unit) {
	c.inner.Incr(c.prefix+m, x, u)
}

func (c prefixedCollector) Set(m string, x any, u units.Unit) {
	c.inner.Set(c.prefix+m, x, u)
}

// func (c prefixedCollector) SetDenom(m string, x any, u units.Unit) {
// 	c.inner.SetDenom(c.prefix+m, x, u)
// }

type contextKey struct{}

func NewContext(ctx context.Context, c Collector) context.Context {
	if c == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, c)
}

func FromContext(ctx context.Context) Collector {
	v := ctx.Value(contextKey{})
	if v == nil {
		return NullCollector{}
	}
	c := v.(Collector)
	if c == nil {
		return NullCollector{}
	}
	return v.(Collector)
}

func Group(ctx context.Context, name string) context.Context {
	if name == "" {
		return ctx
	}
	return NewContext(ctx, prefixedCollector{
		inner:  FromContext(ctx),
		prefix: "/" + name,
	})
}

// Drop removes telemetry from the Context
func Drop(ctx context.Context) context.Context {
	var x Collector
	return context.WithValue(ctx, contextKey{}, x)
}
