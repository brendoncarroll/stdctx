package logctx

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level is the level to log at
type Level = zapcore.Level

const (
	LevelTrace = zapcore.DebugLevel
	LevelDebug = zapcore.DebugLevel

	LevelInfo = zapcore.InfoLevel

	LevelWarn  = zapcore.WarnLevel
	LevelError = zapcore.ErrorLevel

	LevelPanic = zapcore.PanicLevel
	LevelFatal = zapcore.FatalLevel
)

type Attr = zap.Field

type Logger = *zap.Logger

var noOpLogger = zap.New(nil)

type contextKey struct{}

// NewContext returns a new context with sl added to it.
func NewContext(ctx context.Context, l Logger) context.Context {
	if l == nil {
		return ctx
	}
	return context.WithValue(ctx, contextKey{}, l)
}

// FromContext gets an slog.Logger from the context.
// FromContext never returns nil, and defaults to a No-Op logger if one has not been set on the context.
func FromContext(ctx context.Context) Logger {
	v := ctx.Value(contextKey{})
	if v == nil {
		return noOpLogger
	}
	l := v.(Logger)
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
	l := FromContext(ctx).Named(name)
	return NewContext(ctx, l)
}

// Drop removes the logger from the context if it exists
func Drop(ctx context.Context) context.Context {
	var x *Logger
	return context.WithValue(ctx, contextKey{}, x)
}

func Log(ctx context.Context, level Level, msg string, attrs ...Attr) {
	l := FromContext(ctx)
	l.Log(level, msg, attrs...)
}

func Trace(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelTrace, msg, attrs...)
}

func Debug(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelDebug, msg, attrs...)
}

func Info(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelInfo, msg, attrs...)
}

func Warn(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelWarn, msg, attrs...)
}

func Error(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelError, msg, attrs...)
}

func Fatal(ctx context.Context, msg string, attrs ...Attr) {
	Log(ctx, LevelFatal, msg, attrs...)
}

////
// Println based
////

func Logln(ctx context.Context, level Level, args ...any) {
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

func Logf(ctx context.Context, level Level, fmtStr string, args ...any) {
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
