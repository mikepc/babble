package babble

import (
	"fmt"
	"math/rand"
	"strconv"
)

type wordIndexByLength map[string][]string

func addByLength(s string, min int, max int, inMap wordIndexByLength ) {
	slen := len(s)
	if slen >= min && slen <= max {
		l := fmt.Sprintf("%d", slen)
		if _, ok := inMap[l]; ok {
			inMap[l] = append(inMap[l],s )
		}else{
			inMap[l] = []string{ s }
			addByLength(s, min, max, inMap)
		}
	}

}

func sliceWordList(words []string, min int, max int) wordIndexByLength {
	if min < 1 || min > max {
		min = 1
	}
	if max > 36 || min > max {
		max = 36
	}
	d := map[string][]string {}

	for _, s := range words {
		addByLength(s, min, max, d )
	}
	return d
}

func getRandomWordList(min int, max int, list wordIndexByLength) []string {
	keys := []int{}
	for l := range list {
		i, _ := strconv.Atoi(l)
		if  i > max {
			break
		}
		if i >= min {
			keys = append(keys)
		}
	}
	theLength := keys[rand.Int()%len(keys)]
	i := string(theLength)
	a := list[i]
	// return a[rand.Int()%len(a)]
	return a
}
func GenerateEligibleWordList(a []string, min int, max int) []string {
	words := sliceWordList(a, min, max)
	r := []string{}
	for _, a := range words {
		r = append(r, a...)
	}
	return r
}
func getRandomWordFromList(min int, max int, wordList []string) string {
	if len(wordList) == 0{
		panic("word list is not available")
	}
	return wordList[rand.Int()%len(wordList)]
}