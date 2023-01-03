package infrastructure_test

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewTracer", func() {
	It("returns a non-nil ITracer interface", func() {
		env := &lib.Env{
			ServiceName: "test-service",
		}
		Expect(infrastructure.NewTracer(env)).NotTo(BeNil())
	})
})
