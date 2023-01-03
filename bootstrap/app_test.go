package bootstrap_test

import (
	"core-gin/bootstrap"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewApp", func() {
	It("returns an App instance with the correct command", func() {
		// Call the NewApp function
		app := bootstrap.NewApp()

		// Assert that the returned value is an App instance with the correct command
		Expect(app.Command).To(Equal(bootstrap.RootCmd))
	})
})
