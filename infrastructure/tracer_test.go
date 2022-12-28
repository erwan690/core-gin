package infrastructure

import (
	"testing"

	"core-gin/lib"

	"github.com/stretchr/testify/assert"
)

func TestNewTracer(t *testing.T) {
	env := &lib.Env{ServiceName: "test-service"}
	tracer := NewTracer(env)
	assert.NotNil(t, tracer)
}
