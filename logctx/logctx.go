package logctx

import (
	"context"
)

type FmtLogger interface {
	Tracef(fmt string, args ...any)
	Debugf(fmt string, args ...any)
	Infof(fmt string, args ...any)
	Warnf(fmt string, args ...any)
	Errorf(fmt string, args ...any)
	Fatalf(fmt string, args ...any)
}

type logCtxKey struct{}

func WithFmtLogger(ctx context.Context, l FmtLogger) context.Context {
	return context.WithValue(ctx, logCtxKey{}, l)
}

func GetFmtLogger(ctx context.Context) FmtLogger {
	v := ctx.Value(logCtxKey{})
	if v == nil {
		return nil
	}
	return v.(FmtLogger)
}

func Tracef(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Tracef(fmt, args...)
	}
}

func Debugf(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Debugf(fmt, args...)
	}
}

func Infof(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Infof(fmt, args...)
	}
}

func Warnf(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Warnf(fmt, args...)
	}
}

func Errorf(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Errorf(fmt, args...)
	}
}

func Fatalf(ctx context.Context, fmt string, args ...any) {
	if l := GetFmtLogger(ctx); l != nil {
		l.Fatalf(fmt, args...)
	}
}
