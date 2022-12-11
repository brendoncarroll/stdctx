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

type (
	dialerCtxKey   struct{}
	listenerCtxKey struct{}
)

func WithDialer(ctx context.Context, d Dialer) context.Context {
	return context.WithValue(ctx, dialerCtxKey{}, d)
}

func WithoutDialer(ctx context.Context) context.Context {
	return context.WithValue(ctx, dialerCtxKey{}, nil)
}

func GetDialer(ctx context.Context) Dialer {
	v := ctx.Value(dialerCtxKey{})
	if v == nil {
		return nil
	}
	return v.(Dialer)
}

type Listener interface {
	Listen(ctx context.Context, network, addr string) (net.Listener, error)
	ListenPacket(ctx context.Context, network, addr string) (net.PacketConn, error)
}

func WithListener(ctx context.Context, l Listener) context.Context {
	return context.WithValue(ctx, listenerCtxKey{}, l)
}

func WithoutListener(ctx context.Context) context.Context {
	return context.WithValue(ctx, dialerCtxKey{}, nil)
}

func GetListener(ctx context.Context) Listener {
	l := ctx.Value(listenerCtxKey{})
	if l == nil {
		return nil
	}
	return l.(Listener)
}

func Dial(ctx context.Context, network, addr string) (net.Conn, error) {
	d := GetDialer(ctx)
	if d == nil {
		return nil, errors.New("no dialer available in this context")
	}
	return d.DialContext(ctx, network, addr)
}

func Listen(ctx context.Context, network, addr string) (net.Listener, error) {
	l := GetListener(ctx)
	if l == nil {
		return nil, errors.New("no listener availabel in this context")
	}
	return l.Listen(ctx, network, addr)
}

func ListenPacket(ctx context.Context, network, addr string) (net.PacketConn, error) {
	l := GetListener(ctx)
	if l == nil {
		return nil, errors.New("no listener availabe in this context")
	}
	return l.ListenPacket(ctx, network, addr)
}
