package logctx

import (
	"context"
	"fmt"

	"golang.org/x/exp/slog"
)

type Level = slog.Level

const (
	TraceLevel = Level(-8)
	DebugLevel = slog.DebugLevel
	InfoLevel  = slog.InfoLevel
	WarnLevel  = slog.WarnLevel
	ErrorLevel = slog.ErrorLevel
	FatalLevel = Level(16)
)

// NewContext returns a new context with sl added to it.
// It is equivalent to slog.NewContext
func NewContext(ctx context.Context, sl slog.Logger) context.Context {
	return slog.NewContext(ctx, sl)
}

// FromContext gets an slog.Logger from the context.
// It is equivalent to slog.Ctx
func FromContext(ctx context.Context) slog.Logger {
	return slog.Ctx(ctx)
}

func With(ctx context.Context, args ...any) slog.Logger {
	return FromContext(ctx).With(args)
}

func WithGroup(ctx context.Context, name string) slog.Logger {
	return FromContext(ctx).WithGroup(name)
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
