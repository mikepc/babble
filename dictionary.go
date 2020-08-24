package babble

import "strings"

type ExcludeWord func(s string)bool
type TransformWord func(s string) string

type DictionaryConfig struct {
	MinLength int
	MaxLength int
	CustomWordList *[]string
	Downcase bool
	Upcase bool
	ExcludeWord ExcludeWord
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
func (d *babbleDictionary) GetRandomWord() string {
	if d.config.MinLength > d.config.MaxLength {
		panic("minimum length cannot exceed maximum length")
	}
  return getRandomWordFromList(d.config.MinLength, d.config.MaxLength,d.sourceList)
}
func (d *babbleDictionary) GetListLength() int {
	return len(d.sourceList)
}
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
func (d *babbleDictionary) GetWordList() []string {
	return d.sourceList
}
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

