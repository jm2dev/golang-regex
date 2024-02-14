package regex_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"regexp"
)

var _ = Describe("Phonenumber", Label("phonenumbers"), func() {
		r, _ := regexp.Compile("[0-9\\-]{9,11}")

		Context("extract Spanish phone numbers from string", func() {
			It("should find 123456789", func() {
				actual := r.FindString("ejemplo con 123456789 y cosas")
				Expect(actual).To(Equal("123456789"))
			})

			It("should find 123-456-789", func() {
				actual := r.FindString("ejemplo con 123-456-789 y cosas")
				Expect(actual).To(Equal("123-456-789"))
			})
		})

})
