package main

import (
	"strings"
	"testing"
    "github.com/stretchr/testify/assert"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCounts := make(map[string]int)

	for _, v := range words {
		if _, ok := wordCounts[v]; ok {
			wordCounts[v] += 1
		} else {
			wordCounts[v] = 1
		}
	}

	return wordCounts
}

func TestWordCount(t *testing.T) {
    assert := assert.New(t)

    test1 := "I am learning Go!"
    map1 := map[string]int{"Go!":1, "I":1, "am":1, "learning":1}

    test2 := "The quick brown fox jumped over the lazy dog."
    map2 := map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}

    test3 := "I ate a donut. Then I ate another donut."
    map3 := map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2} 

    test4 := "A man a plan a canal panama."
    map4 := map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}

    assert.Equal(map1, WordCount(test1))
    assert.Equal(map2, WordCount(test2))
    assert.Equal(map3, WordCount(test3))
    assert.Equal(map4, WordCount(test4))
}

func main() {

}
