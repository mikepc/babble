Babble
=========

Babble is a small utility that generates random words for you. I found this useful because occasionally you need a random word for testing.

![tower of babel](http://image.shutterstock.com/display_pic_with_logo/518173/140700250/stock-photo-tower-of-babel-first-variant-raster-variant-140700250.jpg)

Usage
-----

```go
package your_app

import (
  "github.com/tjarratt/babble"
)

func main() {
  babbler := babble.NewBabbler()
  println(babbler.Babble()) // excessive-yak-shaving (or some other phrase)

  // optionally set your own separator
  babbler.Separator = " "
  println(babbler.Babble()) // "hello from nowhere" (or some other phrase)

  // optionally set the number of words you want
  babbler.Count = 1
  println(babbler.Babble()) // antibiomicrobrial (or some other word)

  return
})
```
Custom Usage
----
```go
package your_app

import (
  "github.com/tjarratt/babble"
)

func main() {
    c := babble.DictionaryConfig {
				MinLength: 3,
				MaxLength: 5,
			}

  babbler := babble.NewBabblerWithConfig(c)
  babbler.Separator = "-"
  println(babbler.Babble()) // excessive-yak-shaving (or some other phrase)

  // optionally set your own separator
  babbler.Separator = " "
  println(babbler.Babble()) // "hello from nowhere" (or some other phrase)

  // optionally set the number of words you want
  babbler.Count = 1
  println(babbler.Babble()) // antibiomicrobrial (or some other word)

```
Upcase
----
```go
package your_app

import (
"github.com/tjarratt/babble"
)

func main() {
   c := babble.DictionaryConfig {
				MinLength: 3,
				MaxLength: 5,
				Upcase: true,
			}

  babbler := babble.NewBabblerWithConfig(c)
  babbler.Separator = "-"
  println(babbler.Babble()) // EXCESSIVE-YAK-SHAVING (or some other phrase)
```
Downcase
----
```go
  c = babble.DictionaryConfig {
				MinLength: 3,
				MaxLength: 5,
				Downcase: true,
			}

  babbler = babble.NewBabblerWithConfig(c)
  babbler.Separator = "-"
  println(babbler.Babble()) // excessive-yak-shaving (or some other phrase)
```
Transform
----
```go
  var alternateCase = func (s string)string {
                    rs, upper := []rune(s), false
                for i, r := range rs {
                    if unicode.IsLetter(r) {
                        if upper = !upper; upper {
                            rs[i] = unicode.ToUpper(r)
                        }
                    }
                }
                return string(rs)
  }  
  c = babble.DictionaryConfig {
				MinLength: 3,
				MaxLength: 5,
				TransformWord: alternateCase,
			}

  babbler = babble.NewBabblerWithConfig(c)
  babbler.Separator = "-"
  println(babbler.Babble()) // ExCeSsIvE-YaK-SlUrPiNg the transform is applied to each word individually during load
```
Custom Word List
----
```go
customList = []string{ "luke", "leia", "han", "darth", "r2d2" } // C-3PO has been expressly excluded 
 c = babble.DictionaryConfig {
				MinLength: 3,
				MaxLength: 5,
				CustomWordList: customList,
			}

  babbler = babble.NewBabblerWithConfig(c)
  babbler.Separator = "-"
  println(babbler.Babble()) // luke-darth
```

