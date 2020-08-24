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
	b.Dictionary = NewDictionaryWithConfig(DefaultDictionaryConfig)
	b.Count = 2
	b.Separator = "-"
	b.Words = words
	return
}
func NewBabblerWithConfig(config DictionaryConfig) (b Babbler) {
	b.Dictionary = NewDictionaryWithConfig(config)
	b.Count = 2
	b.Separator = "-"
	b.Words = words
	return
}

func (this Babbler) Babble() string {
	pieces := []string{}
	for i := 0; i < this.Count ; i++ {
		pieces = append(pieces, this.Words[rand.Int()%len(this.Words)])
	}
	return strings.Join(pieces, this.Separator)
}
