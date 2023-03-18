package logctx

import (
	"go.uber.org/zap"
)

func String(k string, v string) Attr {
	return zap.String(k, v)
}

func Strings(k string, v []string) Attr {
	return zap.Strings(k, v)
}

func Any(k string, v any) Attr {
	return zap.Any(k, v)
}

func Int(k string, v int) Attr {
	return zap.Int(k, v)
}

func Bool(k string, v bool) Attr {
	return zap.Bool(k, v)
}

func Float64(k string, v float64) Attr {
	return zap.Float64(k, v)
}
