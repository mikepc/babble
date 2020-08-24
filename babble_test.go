package babble_test

import (
	. "babble"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("babble", func() {
	var babbler Babbler
	BeforeEach(func() {
		babbler = Babbler{
			Count: 1,
			Words: []string{"hello"},
			Separator: "☃",
		}
	})

	It("returns a random word", func() {
		Expect(babbler.Babble()).To(Equal("hello"))
	})

	Describe("with multiple words", func() {
		It("concatenates strings", func() {
			babbler.Count = 2
			Expect(babbler.Babble()).To(Equal("hello☃hello"))
		})
	})

	Describe("new babble", func(){
		It("handles a basic configuration", func(){
			babbler = NewBabbler()
			Expect( len(babbler.Words) ).Should(Equal(babbler.Dictionary.GetListLength()))
		})
		It("accepts a custom dictionary config", func(){
			babbler = NewBabblerWithConfig(DictionaryConfig{ MinLength: 2, MaxLength: 4})
			Expect(len(babbler.Words)).Should(Equal(babbler.Dictionary.GetListLength()))
		})

	})
})
