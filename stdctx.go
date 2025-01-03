package stdctx

import (
	"context"

	"go.brendoncarroll.net/stdctx/logctx"
	"go.brendoncarroll.net/stdctx/telctx"
)

type childConfig struct{}

type ChildOption = func(c *childConfig)

// Child creates a child context using name
// Child will use the other libraries in this package to create namespaced resources.
//
// The following resources are managed
// - logctx
// - telctx
func Child(ctx context.Context, name string, opts ...ChildOption) context.Context {
	ctx = logctx.Group(ctx, name)
	ctx = telctx.Group(ctx, name)
	return ctx
}
