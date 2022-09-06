package netctx

import (
	"context"
	"errors"
	"net"
)

type Stack struct {
	net.Dialer
	net.ListenConfig
}

type Dialer interface {
	DialContext(ctx context.Context, netw, addr string) (net.Conn, error)
}

type dialerCtxKey struct{}

func WithDialer(ctx context.Context, d Dialer) context.Context {
	return context.WithValue(ctx, dialerCtxKey{}, d)
}

func WithoutDialer(ctx context.Context) context.Context {
	return context.WithValue(ctx, dialerCtxKey{}, nil)
}

func GetDialer(ctx context.Context) Dialer {
	d := ctx.Value(dialerCtxKey{})
	if d == nil {
		return nil
	}
	return d.(Dialer)
}

func Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	d := GetDialer(ctx)
	if d == nil {
		return nil, errors.New("not dialer available in this context")
	}
	return d.DialContext(ctx, network, addr)
}
