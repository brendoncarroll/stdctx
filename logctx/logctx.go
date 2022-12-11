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
	TraceLevel = Level(-8)
	DebugLevel = slog.DebugLevel
	InfoLevel  = slog.InfoLevel
	WarnLevel  = slog.WarnLevel
	ErrorLevel = slog.ErrorLevel
	FatalLevel = Level(16)
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
		return &noOpLogger
	}
	l := v.(*slog.Logger)
	if l == nil {
		return &noOpLogger
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
	return NewContext(ctx, &l)
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
	Log(ctx, TraceLevel, msg, args)
}

func Debug(ctx context.Context, msg string, args ...any) {
	Log(ctx, DebugLevel, msg, args)
}

func Info(ctx context.Context, msg string, args ...any) {
	Log(ctx, InfoLevel, msg, args)
}

func Warn(ctx context.Context, msg string, args ...any) {
	Log(ctx, WarnLevel, msg, args)
}

func Error(ctx context.Context, msg string, args ...any) {
	Log(ctx, ErrorLevel, msg, args)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	Log(ctx, FatalLevel, msg, args)
}

////
// Println based
////

func Logln(ctx context.Context, level slog.Level, args ...any) {
	sl := FromContext(ctx)
	sl.Log(level, fmt.Sprint(args...))
}

func Traceln(ctx context.Context, args ...any) {
	Logln(ctx, TraceLevel, args...)
}

func Debugln(ctx context.Context, args ...any) {
	Logln(ctx, DebugLevel, args...)
}

func Infoln(ctx context.Context, args ...any) {
	Logln(ctx, InfoLevel, args...)
}

func Warnln(ctx context.Context, args ...any) {
	Logln(ctx, WarnLevel, args...)
}

func Errorln(ctx context.Context, args ...any) {
	Logln(ctx, ErrorLevel, args...)
}

func Fatalln(ctx context.Context, args ...any) {
	Logln(ctx, FatalLevel, args...)
}

////
// Printf based
////

func Logf(ctx context.Context, level slog.Level, fmtStr string, args ...any) {
	sl := FromContext(ctx)
	sl.Log(level, fmt.Sprintf(fmtStr, args...))
}

func Tracef(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, TraceLevel, fmt, args...)
}

func Debugf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, DebugLevel, fmt, args...)
}

func Infof(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, InfoLevel, fmt, args...)
}

func Warnf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, WarnLevel, fmt, args...)
}

func Errorf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, ErrorLevel, fmt, args...)
}

func Fatalf(ctx context.Context, fmt string, args ...any) {
	Logf(ctx, FatalLevel, fmt, args...)
}
