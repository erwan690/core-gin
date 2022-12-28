package infrastructure

import (
	"testing"

	"core-gin/lib"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	env := &lib.Env{Environment: "development", MaxMultipartMemory: 10485760}
	router := NewRouter(env)
	assert.Equal(t, gin.DebugMode, gin.Mode())
	assert.Equal(t, true, router.Engine.RemoveExtraSlash)
	assert.Equal(t, int64(10485760), router.Engine.MaxMultipartMemory)

	env = &lib.Env{Environment: "production", MaxMultipartMemory: 10485760}
	router = NewRouter(env)
	assert.Equal(t, gin.ReleaseMode, gin.Mode())
	assert.Equal(t, true, router.Engine.RemoveExtraSlash)
	assert.Equal(t, int64(10485760), router.Engine.MaxMultipartMemory)
}
