package utils_test

import (
	"core-gin/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("StatusInList", func() {
		It("returns true when the status is in the list", func() {
			result := utils.StatusInList(200, []int{200, 404, 500})
			Expect(result).To(BeTrue())
		})

		It("returns false when the status is not in the list", func() {
			result := utils.StatusInList(301, []int{200, 404, 500})
			Expect(result).To(BeFalse())
		})

		It("returns false when the list is empty", func() {
			result := utils.StatusInList(200, []int{})
			Expect(result).To(BeFalse())
		})
	})
})
