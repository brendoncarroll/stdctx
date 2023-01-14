package logctx

import (
	"context"
	"fmt"
	"io"

	"golang.org/x/exp/slog"
)

type Logger = slog.Logger

type Level = slog.Level

const (
	LevelTrace = Level(-8)
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = Level(16)
)

type contextKey struct{}

// NewContext returns a new context with sl added to it.
func NewContext(ctx context.Context, sl *Logger) context.Context {
	if sl == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, sl)
}

var noOpLogger = slog.New(slog.NewTextHandler(io.Discard))

// FromContext gets an slog.Logger from the context.
// FromContext never returns nil, and defaults to a No-Op logger if one has not been set on the context.
func FromContext(ctx context.Context) *slog.Logger {
	v := ctx.Value(contextKey{})
	if v == nil {
		return noOpLogger
	}
	l := v.(*slog.Logger)
	if l == nil {
		return noOpLogger
	}
	return l
}

// IsSet returns true if a logger is set on the context
func IsSet(ctx context.Context) bool {
	v := ctx.Value(contextKey{})
	return v != nil && v.(*Logger) != nil
}

func Group(ctx context.Context, name string) context.Context {
	l := FromContext(ctx).WithGroup(name)
	return NewContext(ctx, l)
}

// Drop removes the logger from the context if it exists
func Drop(ctx context.Context) context.Context {
	var x *Logger
	return context.WithValue(ctx, contextKey{}, x)
}

func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	sl := FromContext(ctx)
	sl.Log(level, msg, args...)
}

func Trace(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelTrace, msg, args)
}

func Debug(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelDebug, msg, args)
}

func Info(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelInfo, msg, args)
}

func Warn(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelWarn, msg, args)
}

func Error(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelError, msg, args)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	Log(ctx, LevelFatal, msg, args)
}

////
// Println based
////

func Logln(ctx context.Context, level slog.Level, args ...any) {
	sl := FromContext(ctx)
	sl.Log(level, fmt.Sprint(args...))
}

func Traceln(ctx context.Context, args ...any) {
	Logln(ctx, LevelTrace, args...)
}

func Debugln(ctx context.Context, args ...any) {
	Logln(ctx, LevelDebug, args...)
}

func Infoln(ctx context.Context, args ...any) {
	Logln(ctx, LevelInfo, args...)
}

func Warnln(ctx context.Context, args ...any) {
	Logln(ctx, LevelWarn, args...)
}

func Errorln(ctx context.Context, args ...any) {
	Logln(ctx, LevelError, args...)
}

func Fatalln(ctx context.Context, args ...any) {
	Logln(ctx, LevelFatal, args...)
}

////
// Printf based
////

func Logf(ctx context.Context, level slog.Level, fmtStr string, args ...any) {
	sl := FromContext(ctx)
	sl.Log(level, fmt.Sprintf(fmtStr, args...))
}

func Tracef(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelTrace, fmt, args...)
}

func Debugf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelDebug, fmt, args...)
}

func Infof(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelInfo, fmt, args...)
}

func Warnf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelWarn, fmt, args...)
}

func Errorf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelError, fmt, args...)
}

func Fatalf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, LevelFatal, fmt, args...)
}
