package babble_test

import (
	"github.com/mikepc/babble"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("dictionary", func() {
	var d babble.Dictionary
	var w string
	BeforeEach(func() {
		d = babble.NewDictionaryWithConfig(babble.DictionaryConfig{ MinLength: 3, MaxLength: 5 })
		w = d.GetRandomWord()
	})

	It("returns a random word", func() {
		s := d.GetRandomWord()
		ls := len(s)
		Expect(ls >= 1).Should(BeTrue())
		Expect(ls <= 5).Should(BeTrue())

	})

	Describe("with custom configuration", func() {
		It("handles min 0 max 30", func() {
			d = babble.NewDictionaryWithConfig(babble.DictionaryConfig{MinLength: 0, MaxLength: 30})
			s := d.GetRandomWord()
			ls := len(s)
			Expect(ls >= 0).Should(BeTrue())
			Expect(ls <= 30).Should(BeTrue())
		})

		It("handles min 3 max 3", func() {
			d = babble.NewDictionaryWithConfig(babble.DictionaryConfig{ MinLength: 3, MaxLength: 3})
			s := d.GetRandomWord()
			ls := len(s)
			Expect(ls >= 3).Should(BeTrue())
			Expect(ls <= 3).Should(BeTrue())
		})

		It("handles when min length is larger than max, should use panic", func() {
			d = babble.NewDictionaryWithConfig(babble.DictionaryConfig{ MinLength: 30, MaxLength: 3})
			defer func(){
				 r := recover();
				 Expect(r).ShouldNot(BeNil())
			}()
			d.GetRandomWord()

		})
		It("ExcludeWord", func() {
			Expect(w).ShouldNot(BeEmpty())
			a := d.GetListLength()
			exc := func(s string)bool {
				return s == w
			}
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				ExcludeWord: exc,
			}
			d =  babble.NewDictionaryWithConfig(c)
			b := d.GetListLength()
			Expect(b == a - 1).Should(BeTrue())
		})
		It("TransformWord", func() {
			Expect(w).ShouldNot(BeEmpty())
			tw := func(s string)string {
				r := []rune(s)
				for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
					r[i], r[j] = r[j], r[i]
				}
				return string(r)
			}
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				TransformWord: tw,
			}
			d =  babble.NewDictionaryWithConfig(c)
			aw := tw(w) // transform the test word
			Expect(aw).ShouldNot(Equal(w))
			words := d.GetWordList()
			found := false
			for _, txwd := range words {
				if txwd == aw {
					found = true
				}
			}
			Expect(found).Should(BeTrue())

		})
		It("uses a custom word list", func(){
			Expect(w).ShouldNot(BeEmpty())
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				CustomWordList: &[]string{ w },
			}
			d =  babble.NewDictionaryWithConfig(c)
			actual := d.GetRandomWord()
			Expect(actual).Should(Equal(w))
		})
		It("upcases",  func() {
			Expect(w).ShouldNot(BeEmpty())
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				Upcase: true,
				CustomWordList: &[]string{ w },
			}
			d =  babble.NewDictionaryWithConfig(c)
			actual := d.GetRandomWord()
			Expect(actual).Should(Equal(strings.ToUpper(w)))


		})
		It("downcases",  func() {
			Expect(w).ShouldNot(BeEmpty())
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				Downcase: true,
				CustomWordList: &[]string{ w },
			}
			d =  babble.NewDictionaryWithConfig(c)
			actual := d.GetRandomWord()
			Expect(actual).Should(Equal(strings.ToLower(w)))
		})
		It("cannot upcase and downcase",  func() {
			defer func(){
				r := recover();
				Expect(r).ShouldNot(BeNil())
			}()
			c := babble.DictionaryConfig{
				MinLength: 3,
				MaxLength: 5,
				Downcase: true,
				Upcase: true,
				CustomWordList: &[]string{ w },
			}
			d =  babble.NewDictionaryWithConfig(c)

		})
	})
})

