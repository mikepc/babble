package babble

import (
	"math/rand"
	"strings"
	"time"
)

var words []string

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Babbler struct {
	Count int
	Separator string
	Words []string
	Dictionary Dictionary
}

func NewBabbler() (b Babbler) {
	d := NewDictionaryWithConfig(DefaultDictionaryConfig)
	b.Dictionary = d
	b.Count = 2
	b.Separator = "-"
	b.Words = d.GetWordList()
	return
}
func NewBabblerWithConfig(config DictionaryConfig) (b Babbler) {
	d := NewDictionaryWithConfig(config)
	b.Dictionary = d
	b.Count = 2
	b.Separator = "-"
	b.Words = b.Dictionary.GetWordList()
	return
}

func (this Babbler) Babble() string {
	pieces := []string{}
	for i := 0; i < this.Count ; i++ {
		pieces = append(pieces, this.Words[rand.Int()%len(this.Words)])
	}
	return strings.Join(pieces, this.Separator)
}
