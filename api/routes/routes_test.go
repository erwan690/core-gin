package routes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "core-gin/api/routes"
)

type MockRoute struct {
	SetupCalled bool
}

func (m *MockRoute) Setup() {
	m.SetupCalled = true
}

var _ = Describe("Routes", func() {
	var (
		mockHealthRoute *MockRoute
		routes          Routes
	)

	BeforeEach(func() {
		mockHealthRoute = new(MockRoute)
		routes = NewRoutes(mockHealthRoute)
	})

	It("creates a Routes instance with the correct route", func() {
		Expect(routes).To(HaveLen(1))
		Expect(routes[0]).To(Equal(mockHealthRoute))
	})

	It("calls the Setup method on the mock health route", func() {
		routes.Setup()
		Expect(mockHealthRoute.SetupCalled).To(BeTrue())
	})
})
