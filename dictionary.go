package babble

import "strings"

type ExcludeWord func(s string)bool
type TransformWord func(s string) string

type DictionaryConfig struct {
	// Limits the lower bound to the length of words in the list
	MinLength int
	// Limits the upper bound to the length of words in the list
	MaxLength int
	// If this pointer is non-nil, this list is used instead of
	// the default OS dictionary.
	CustomWordList *[]string
	// Changes all letters in each word to lower case
	Downcase bool
	// Changes all letters in each word to upper case
	Upcase bool
	// If present, this function is emitted on all words in the dictionary,
	// if the provided function returns true, the word is not included in the
	// word list.
	ExcludeWord ExcludeWord
	// Emittted on all words in the dictionary
	TransformWord TransformWord
}
var  DefaultDictionaryConfig = DictionaryConfig {
	MinLength: 3,
	MaxLength: 6,
}

type Dictionary interface {
	GetRandomWord() string
	GetListLength() int
	GetWordList() []string
}

type babbleDictionary struct {
	sourceList []string
	config *DictionaryConfig
}
// NewDictionaryWithConfig creates a new dictionary with the supplied
// configuration struct.
//    c = babble.DictionaryConfig {
//				MinLength: 3,
//				MaxLength: 5,
//				Downcase: true,
//
//        }
//    d = NewDictionaryWithConfig(c)
func NewDictionaryWithConfig(c DictionaryConfig) Dictionary {
	if c.Upcase && c.Downcase {
		panic("cannot upcase and downcase the dictionary at the same time, invalid dictionary config")
	}
	d := &babbleDictionary{
		config: &c,
		sourceList: []string{},
	}
	if c.CustomWordList == nil {
		d.load()
	}else{
		d.sourceList = *c.CustomWordList
		if d.config.ExcludeWord != nil {
			d.loadWithExclude()
		} else {
			newSource := []string{}
			for _, s := range d.sourceList {
				newSource = append(newSource, d.applyTransform(s))
			}
			d.sourceList = newSource
		}
	}
	return d
}

// GetRandomWord returns a random word from the dictionary
func (d *babbleDictionary) GetRandomWord() string {
	if d.config.MinLength > d.config.MaxLength {
		panic("minimum length cannot exceed maximum length")
	}
  return getRandomWordFromList(d.config.MinLength, d.config.MaxLength,d.sourceList)
}

// GetListLength returns the length of the currently loaded word list.
func (d *babbleDictionary) GetListLength() int {
	return len(d.sourceList)
}

// applyTransform applies the configured transformation function
// on each string, and if Upcase or Downcase is true, that
// transformation is applied
func (d *babbleDictionary) applyTransform(s string) (ts string) {
	txf := d.config.TransformWord
	if txf != nil {
		ts = txf(s)
	}else{
		ts = s
	}
	if d.config.Downcase {
		return strings.ToLower(ts)
	}
	if d.config.Upcase {
		return strings.ToUpper(ts)
	}
	return ts
}

// GetWordList returns the entire list of possible words
func (d *babbleDictionary) GetWordList() []string {
	return d.sourceList
}

// loadWithExclude loads the list of possible words, emitting
// the ExcludeWord function if present to filter candidate
// words.
func (d *babbleDictionary) loadWithExclude() {
	exc := d.config.ExcludeWord
	newSource := []string{}
	for _, s := range d.sourceList{
		if(!exc(s)){
			newSource = append(newSource, d.applyTransform(s))
		}
	}
	d.sourceList = newSource
}

// load loads the list of possible words
func (d *babbleDictionary) load()  {
	 list := readAvailableDictionary()
	 d.sourceList = GenerateEligibleWordList(list, d.config.MinLength, d.config.MaxLength)
	 if d.config.ExcludeWord != nil {
		 d.loadWithExclude()
	 } else {
	 	newSource := []string{}
		for _, s := range d.sourceList {
			newSource = append(newSource, d.applyTransform(s))
		}
		d.sourceList = newSource
	 }
}

