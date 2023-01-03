package infrastructure_test

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewDatabase", func() {
	It("returns a non-nil Database struct", func() {
		logger := lib.GetLogger()
		env := &lib.Env{
			DBUsername: "username",
			DBPassword: "password",
			DBHost:     "localhost",
			DBPort:     "5432",
			DBName:     "database",
		}
		Expect(infrastructure.NewDatabase(logger, env)).NotTo(BeNil())
	})
})
