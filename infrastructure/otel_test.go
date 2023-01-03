package infrastructure_test

import (
	"core-gin/lib"

	"core-gin/infrastructure"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewOtel", func() {
	It("returns a non-nil Otel struct", func() {
		logger := lib.GetLogger()
		env := &lib.Env{
			OtelEnable:   true,
			OtelEndpoint: "localhost:55680",
			ServiceName:  "test-service",
		}
		Expect(infrastructure.NewOtel(env, logger)).NotTo(BeNil())
	})

	It("returns a non-nil Otel struct with insecure", func() {
		logger := lib.GetLogger()
		env := &lib.Env{
			InsecureMode: true,
			OtelEnable:   true,
			OtelEndpoint: "localhost:55680",
			ServiceName:  "test-service",
		}
		Expect(infrastructure.NewOtel(env, logger)).NotTo(BeNil())
	})

	It("returns a nil Otel struct with disable", func() {
		logger := lib.GetLogger()
		env := &lib.Env{
			InsecureMode: true,
			OtelEnable:   false,
			OtelEndpoint: "localhost:55680",
			ServiceName:  "test-service",
		}
		Expect(infrastructure.NewOtel(env, logger)).To(Equal(infrastructure.Otel{
			Exporter: nil,
		}))
	})
})
