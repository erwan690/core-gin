package lib

import "core-gin/utils"

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	ServiceName string `mapstructure:"SERVICE_NAME"`
	// Otel Config
	OtelEnable   bool   `mapstructure:"OTEL_ENABLE"`
	InsecureMode bool   `mapstructure:"OTEL_INSECURE_MODE"`
	OtelEndpoint string `mapstructure:"OTEL_ENDPOINT"`

	// Otel Config
	RateLimitPeriod   int64 `mapstructure:"RATE_LIMIT_PEROID"`
	RateLimitRequests int64 `mapstructure:"RATE_LIMIT_REQUEST"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	MaxMultipartMemory int64 `mapstructure:"MAX_MULTIPART_MEMORY"`

	TimeZone  string `mapstructure:"TIMEZONE"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() *Env {
	globalEnv.LogLevel = utils.GetEnv("LOG_LEVEL", "info")
	globalEnv.ServerPort = utils.GetEnv("SERVER_PORT", "4040")
	globalEnv.Environment = utils.GetEnv("ENVIRONMENT", "local")

	globalEnv.ServiceName = utils.GetEnv("SERVICE_NAME", "core-gin")

	globalEnv.OtelEnable = utils.GetEnvAsBool("OTEL_ENABLE", false)
	globalEnv.InsecureMode = utils.GetEnvAsBool("OTEL_INSECURE_MODE", true)
	globalEnv.OtelEndpoint = utils.GetEnv("OTEL_ENDPOINT", "local")

	globalEnv.RateLimitPeriod = utils.GetEnvAsInt64("RATE_LIMIT_PEROID", 5)
	globalEnv.RateLimitRequests = utils.GetEnvAsInt64("RATE_LIMIT_REQUEST", 200)

	globalEnv.DBUsername = utils.GetEnv("DB_USER", "")
	globalEnv.DBPassword = utils.GetEnv("DB_PASS", "")
	globalEnv.DBHost = utils.GetEnv("DB_HOST", "")
	globalEnv.DBPort = utils.GetEnv("DB_PORT", "")
	globalEnv.DBName = utils.GetEnv("DB_NAME", "")

	globalEnv.TimeZone = utils.GetEnv("TIMEZONE", "UTC")

	globalEnv.JWTSecret = utils.GetEnv("JWT_SECRET", "")

	return &globalEnv
}
