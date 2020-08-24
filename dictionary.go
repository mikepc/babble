package babble

type DictionaryConfig struct {
	MinLength int
	MaxLength int
	CustomWordList *[]string
}
var  DefaultDictionaryConfig = DictionaryConfig {
	MinLength: 3,
	MaxLength: 6,
}

type Dictionary interface {
	SetExcludeFunc(func(s string)bool) // Matches words in list, strings returning true are excluded
	SetTransformFunc(func(s string)string) // Executes transformation function when building the list
	GetRandomWord() string
	GetListLength() int
}

type babbleDictionary struct {
	sourceList []string
	excludeFunc  func(s string)bool
	transformFunc func(s string)string
	config *DictionaryConfig
}

func NewDictionaryWithConfig(c DictionaryConfig) Dictionary {
	d := &babbleDictionary{
		config: &c,
		sourceList: []string{},

	}
	if c.CustomWordList == nil {
		d.sourceList =	d.load()
	}else{
		d.sourceList = *c.CustomWordList
	}
	return d
}
func (d *babbleDictionary) GetRandomWord() string {
  return getRandomWordFromList(d.config.MinLength, d.config.MaxLength,d.sourceList)
}
func (d *babbleDictionary) GetListLength() int {
	return len(d.sourceList)
}
func (d *babbleDictionary) load() []string {
	 list := readAvailableDictionary()
	 allEligible := GenerateEligibleWordList(list, d.config.MinLength, d.config.MaxLength)
	 f := []string{}
	 if d.excludeFunc != nil {
	   for _, w := range allEligible{
		  if !d.excludeFunc(w) {
			if d.transformFunc != nil {
				w = d.transformFunc(w)
			}
		  }
	   }
	 }else{
		 for _, w := range allEligible {
			 if d.transformFunc != nil {
				 w = d.transformFunc(w)
			 }
			 f = append(f, w)
		 }
	 }
	 return f
}

func (d *babbleDictionary) SetExcludeFunc(f func(s string)bool){
	d.excludeFunc = f
}
func (d *babbleDictionary) SetTransformFunc(f func(s string)string){
	d.transformFunc = f
}