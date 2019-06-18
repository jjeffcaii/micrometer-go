package micrometer_test

import (
	"testing"

	. "github.com/jjeffcaii/micrometer-go"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	reg := NewSimpleMeterRegistry()
	c, err := Builder().
		Counter("rsocket.request").
		Tags("service", "ping.PingPong", "method", "ping").
		Register(reg)
	require.NoError(t, err, "builder failed")
	require.NotNil(t, c, "nil counter")
}
