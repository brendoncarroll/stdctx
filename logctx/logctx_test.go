package logctx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNil(t *testing.T) {
	l := FromContext(context.Background())
	require.NotNil(t, l)
	require.False(t, IsSet(context.Background()))
}
