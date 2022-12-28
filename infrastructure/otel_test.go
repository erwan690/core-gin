package infrastructure

import (
	"testing"

	"core-gin/lib"

	"github.com/stretchr/testify/assert"
)

func TestNewOtel(t *testing.T) {
	env := &lib.Env{OtelEnable: true, OtelEndpoint: "test-endpoint", ServiceName: "test-service", InsecureMode: true}
	otel := NewOtel(env, lib.GetLogger())
	assert.NotNil(t, otel.Exporter)

	env = &lib.Env{OtelEnable: false}
	otel = NewOtel(env, lib.GetLogger())
	assert.Nil(t, otel.Exporter)
}
