package lib

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	// Test with default values
	os.Clearenv()
	env := NewEnv()
	assert.Equal(t, "info", env.LogLevel)
	assert.Equal(t, "4040", env.ServerPort)
	assert.Equal(t, "local", env.Environment)
	assert.Equal(t, "core-gin", env.ServiceName)
	assert.Equal(t, false, env.OtelEnable)
	assert.Equal(t, true, env.InsecureMode)
	assert.Equal(t, "local", env.OtelEndpoint)
	assert.Equal(t, int64(5), env.RateLimitPeriod)
	assert.Equal(t, int64(200), env.RateLimitRequests)
	assert.Equal(t, "", env.DBUsername)
	assert.Equal(t, "", env.DBPassword)
	assert.Equal(t, "", env.DBHost)
	assert.Equal(t, "", env.DBPort)
	assert.Equal(t, "", env.DBName)
	assert.Equal(t, "UTC", env.TimeZone)
	assert.Equal(t, "", env.JWTSecret)

	// Test with custom values
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("ENVIRONMENT", "prod")
	os.Setenv("SERVICE_NAME", "my-service")
	os.Setenv("OTEL_ENABLE", "true")
	os.Setenv("OTEL_INSECURE_MODE", "false")
	os.Setenv("OTEL_ENDPOINT", "remote")
	os.Setenv("RATE_LIMIT_PEROID", "10")
	os.Setenv("RATE_LIMIT_REQUEST", "500")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASS", "pass")
	os.Setenv("DB_HOST", "host")
	os.Setenv("DB_PORT", "port")
	os.Setenv("DB_NAME", "db-test")
	os.Setenv("TIMEZONE", "Asia/Jakarta")
	os.Setenv("JWT_SECRET", "secret")

	loadedenv := NewEnv()
	assert.Equal(t, "debug", loadedenv.LogLevel)
	assert.Equal(t, "8080", loadedenv.ServerPort)
	assert.Equal(t, "prod", loadedenv.Environment)
	assert.Equal(t, "my-service", loadedenv.ServiceName)
	assert.Equal(t, true, loadedenv.OtelEnable)
	assert.Equal(t, false, loadedenv.InsecureMode)
	assert.Equal(t, "remote", loadedenv.OtelEndpoint)
	assert.Equal(t, int64(10), loadedenv.RateLimitPeriod)
	assert.Equal(t, int64(500), loadedenv.RateLimitRequests)
	assert.Equal(t, "user", loadedenv.DBUsername)
	assert.Equal(t, "pass", loadedenv.DBPassword)
	assert.Equal(t, "host", loadedenv.DBHost)
	assert.Equal(t, "port", loadedenv.DBPort)
	assert.Equal(t, "db-test", loadedenv.DBName)
	assert.Equal(t, "Asia/Jakarta", loadedenv.TimeZone)
	assert.Equal(t, "secret", loadedenv.JWTSecret)
}

func TestGetEnv(t *testing.T) {
	env := GetEnv()
	if env != globalEnv {
		t.Errorf("Expected env to be %v, got %v", globalEnv, env)
	}
}
