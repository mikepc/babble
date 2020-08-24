package babble_test

import (
	"babble"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("dictionary", func() {
	var d babble.Dictionary
	BeforeEach(func() {
		d = babble.NewDictionaryWithConfig(babble.DictionaryConfig{ 1, 5, nil})
	})

	It("returns a random word", func() {
		s := d.GetRandomWord()
		ls := len(s)
		Expect(ls >= 1).Should(BeTrue())
		Expect(ls <= 5).Should(BeTrue())

	})

	Describe("with multiple words", func() {
		It("concatenates strings", func() {

		})
	})
})

