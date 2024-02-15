package regex_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"regexp"
)

var _ = Describe("Phone numbers extraction", Label("phonenumbers"), func() {
	r, _ := regexp.Compile("(\\+34 ){0,1}[0-9-]{9,11}")

		When("there is a Spanish phone number in string", func() {
			It("should find 123456789", func() {
				actual := r.FindString("ejemplo con 123456789 y cosas")
				Expect(actual).To(Equal("123456789"))
			})

			It("should find 123-456-789", func() {
				actual := r.FindString("ejemplo con 123-456-789 y cosas")
				Expect(actual).To(Equal("123-456-789"))
			})

			It("should find +34 123456789", func() {
				actual := r.FindString("ejemplo con +34 123456789 y cosas")
				Expect(actual).To(Equal("+34 123456789"))
			})
		})

})
