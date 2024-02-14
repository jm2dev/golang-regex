package regex_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"regexp"
)

var _ = Describe("Challenge 1", Label("urls"), func() {
	Describe("URLs", func() {
		r, _ := regexp.Compile("(https://www.example.[a-z]{3}([/a-z0-9\\-]*))")

		Context("extract url from string", func() {
			It("should find https://www.example.com", func() {
				actual := r.FindString("ejemplo con https://www.example.com y cosas")
				Expect(actual).To(Equal("https://www.example.com"))
			})

			It("should find https://www.example.net", func() {
				actual := r.FindString("ejemplo con https://www.example.net y cosas")
				Expect(actual).To(Equal("https://www.example.net"))
			})

			It("should find https://www.example.net/lorem/ipsum", func() {
				actual := r.FindString("ejemplo con https://www.example.net/lorem/ipsum y cosas")
				Expect(actual).To(Equal("https://www.example.net/lorem/ipsum"))
			})

		})

		Context("extract urls from string", func() {
			It("should find https://www.example.com and https://www.example.net", func() {
				actual := r.FindAllString("ejemplo con https://www.example.com y variante https://www.example.net y cosas", -1)
				expected := []string{"https://www.example.com", "https://www.example.net"}
				Expect(actual).To(Equal(expected))
			})
		})

		Context("challenge 1 input", func() {
			It("should find three urls", func() {
				actual := r.FindAllString("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla facilisis, mauris ac ultrices tincidunt, velit nunc tincidunt nisl, at lacinia justo nisl id nunc. Sed euismod, mauris ac consequat tincidunt, est est luctus tellus, at https://www.example.com/lorem/ipsum/dolor sit amet mauris. Fusce auctor, https://www.example.com/lorem-ipsum/dolor/sit/amet consectetur adipiscing elit, justo nisl aliquam nunc, at consequat nunc nisl id nunc. In hac habitasse platea dictumst. Sed vitae nunc auctor, consectetur elit at, https://www.example.com/lorem/ipsum/dolor/sit/amet/12345/67890", -1)
				expected := []string{"https://www.example.com/lorem/ipsum/dolor", "https://www.example.com/lorem-ipsum/dolor/sit/amet", "https://www.example.com/lorem/ipsum/dolor/sit/amet/12345/67890"}
				Expect(actual).To(Equal(expected))
			})
		})
	})
})
