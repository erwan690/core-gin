package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	ServiceName string `mapstructure:"SERVICE_NAME"`
	// Otel Config
	InsecureMode bool   `mapstructure:"OTEL_INSECURE_MODE"`
	OtelEndpoint string `mapstructure:"OTEL_ENDPOINT"`

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
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read cofiguration", err)
	}

	viper.SetDefault("TIMEZONE", "UTC")

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	return &globalEnv
}
