package infrastructure_test

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewRouter", func() {
	It("sets the correct Gin mode and default values", func() {
		env := &lib.Env{
			Environment:        "development",
			MaxMultipartMemory: 10485760,
		}
		router := infrastructure.NewRouter(env)
		Expect(gin.Mode()).To(Equal(gin.DebugMode))
		Expect(router.Engine.RemoveExtraSlash).To(BeTrue())
		Expect(router.Engine.MaxMultipartMemory).To(Equal(int64(10485760)))
	})
	It("sets the correct Gin mode and default values production", func() {
		env := &lib.Env{
			Environment:        "production",
			MaxMultipartMemory: 10485760,
		}
		router := infrastructure.NewRouter(env)
		Expect(gin.Mode()).To(Equal(gin.ReleaseMode))
		Expect(router.Engine.RemoveExtraSlash).To(BeTrue())
		Expect(router.Engine.MaxMultipartMemory).To(Equal(int64(10485760)))
	})
})
