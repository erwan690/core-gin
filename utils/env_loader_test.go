package utils_test

import (
	"os"

	"core-gin/utils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("GetEnv", func() {
		It("returns the value of the environment variable when it exists", func() {
			os.Setenv("TEST_VAR", "test_value")
			val := utils.GetEnv("TEST_VAR", "default_value")
			Expect(val).To(Equal("test_value"))
		})

		It("returns the default value when the environment variable does not exist", func() {
			val := utils.GetEnv("NON_EXISTENT_VAR", "default_value")
			Expect(val).To(Equal("default_value"))
		})
	})

	Describe("GetEnvAsInt", func() {
		It("returns the value of the environment variable when it exists and is a valid integer", func() {
			os.Setenv("TEST_VAR", "123")
			val := utils.GetEnvAsInt("TEST_VAR", 0)
			Expect(val).To(Equal(123))
		})

		It("returns the default value when the environment variable exists but is not a valid integer", func() {
			os.Setenv("TEST_VAR", "invalid_int")
			val := utils.GetEnvAsInt("TEST_VAR", 0)
			Expect(val).To(Equal(0))
		})

		It("returns the default value when the environment variable does not exist", func() {
			val := utils.GetEnvAsInt("NON_EXISTENT_VAR", 0)
			Expect(val).To(Equal(0))
		})
	})

	Describe("GetEnvAsInt64", func() {
		It("returns the value of the environment variable when it exists and is a valid integer", func() {
			os.Setenv("TEST_VAR", "123")
			val := utils.GetEnvAsInt64("TEST_VAR", 0)
			Expect(val).To(Equal(int64(123)))
		})

		It("returns the default value when the environment variable exists but is not a valid integer", func() {
			os.Setenv("TEST_VAR", "invalid_int")
			val := utils.GetEnvAsInt64("TEST_VAR", 0)
			Expect(val).To(Equal(int64(0)))
		})

		It("returns the default value when the environment variable does not exist", func() {
			val := utils.GetEnvAsInt64("NON_EXISTENT_VAR", 0)
			Expect(val).To(Equal(int64(0)))
		})
	})

	Describe("GetEnvAsBool", func() {
		It("returns the value of the environment variable when it exists and is a valid boolean", func() {
			os.Setenv("TEST_VAR", "true")
			result := utils.GetEnvAsBool("TEST_VAR", false)
			Expect(result).To(BeTrue())

			os.Setenv("TEST_VAR", "false")
			result = utils.GetEnvAsBool("TEST_VAR", true)
			Expect(result).To(BeFalse())
		})

		It("returns the default value when the environment variable exists but is not a valid boolean", func() {
			os.Setenv("TEST_VAR", "invalid")
			result := utils.GetEnvAsBool("TEST_VAR", false)
			Expect(result).To(BeFalse())
		})

		It("returns the default value when the environment variable does not exist", func() {
			os.Unsetenv("TEST_VAR")
			result := utils.GetEnvAsBool("TEST_VAR", false)
			Expect(result).To(BeFalse())
		})
	})
})
