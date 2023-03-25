package logctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNil(t *testing.T) {
	l := FromContext(context.Background())
	require.NotNil(t, l)
	require.False(t, IsSet(context.Background()))
}

func TestLog(t *testing.T) {
	ctx := context.Background()
	l, _ := zap.NewDevelopment()
	ctx = NewContext(ctx, l)
	Traceln(ctx, "traceln")
	Debugln(ctx, "debugln")
	Infoln(ctx, "infoln")
	Warnln(ctx, "warnln")
	Errorln(ctx, "errorln")
}
