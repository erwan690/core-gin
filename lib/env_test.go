package lib_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"core-gin/lib"
)

var _ = Describe("Lib", func() {
	Describe("NewEnv", func() {
		BeforeEach(func() {
			os.Clearenv()
		})

		It("returns a new environment with default values when no environment variables are set", func() {
			env := lib.NewEnv()
			Expect(env.LogLevel).To(Equal("info"))
			Expect(env.ServerPort).To(Equal("4040"))
			Expect(env.Environment).To(Equal("local"))
			Expect(env.ServiceName).To(Equal("core-gin"))
			Expect(env.OtelEnable).To(BeFalse())
			Expect(env.InsecureMode).To(BeTrue())
			Expect(env.OtelEndpoint).To(Equal("local"))
			Expect(env.RateLimitPeriod).To(Equal(int64(5)))
			Expect(env.RateLimitRequests).To(Equal(int64(200)))
			Expect(env.DBUsername).To(Equal(""))
			Expect(env.DBPassword).To(Equal(""))
			Expect(env.DBHost).To(Equal(""))
			Expect(env.DBPort).To(Equal(""))
			Expect(env.DBName).To(Equal(""))
			Expect(env.TimeZone).To(Equal("UTC"))
			Expect(env.JWTSecret).To(Equal(""))
			Expect(env.SlackToken).To(Equal(""))
			Expect(env.SlackMaintener).To(Equal(""))
			Expect(env.SlackChannelID).To(Equal(""))
		})
		It("returns a new environment with environment variables are set", func() {
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
			os.Setenv("SLACK_TOKEN", "test-token")
			os.Setenv("SLACK_MAINTENER", "test-main")
			os.Setenv("SLACK_CHANNEL", "test-chan")
			env := lib.NewEnv()
			Expect(env.LogLevel).To(Equal("debug"))
			Expect(env.ServerPort).To(Equal("8080"))
			Expect(env.Environment).To(Equal("prod"))
			Expect(env.ServiceName).To(Equal("my-service"))
			Expect(env.OtelEnable).To(BeTrue())
			Expect(env.InsecureMode).To(BeFalse())
			Expect(env.OtelEndpoint).To(Equal("remote"))
			Expect(env.RateLimitPeriod).To(Equal(int64(10)))
			Expect(env.RateLimitRequests).To(Equal(int64(500)))
			Expect(env.DBUsername).To(Equal("user"))
			Expect(env.DBPassword).To(Equal("pass"))
			Expect(env.DBHost).To(Equal("host"))
			Expect(env.DBPort).To(Equal("port"))
			Expect(env.DBName).To(Equal("db-test"))
			Expect(env.TimeZone).To(Equal("Asia/Jakarta"))
			Expect(env.JWTSecret).To(Equal("secret"))
			Expect(env.SlackToken).To(Equal("test-token"))
			Expect(env.SlackMaintener).To(Equal("test-main"))
			Expect(env.SlackChannelID).To(Equal("test-chan"))
		})
	})
})
