package text

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type TextAnalyzer struct {
	Words                       []String
	MostFrequentlyWordsQuantity int
}

// Main method of analyzer which will find the most frequently words
func (ta *TextAnalyzer) FindMostFrequentlyWords() {
	ta.prepareTextWords()

	dict := &WordsDictionary{}

	// search most frequently words
	for _, word := range ta.Words {
		dict.AddWord(word)
	}
	sort.Sort(dict)

	mostFrequentWord := make([]Word, 0, ta.MostFrequentlyWordsQuantity)
	for i := 0; i < ta.MostFrequentlyWordsQuantity; i++ {
		mostFrequentWord = append(mostFrequentWord, *dict.GetWordByIndex(i))
	}

	// output
	for _, word := range mostFrequentWord {
		fmt.Printf("Word \"%s\" - %d times\n", word.Text, word.Count)
	}
}

// Preparaions before counting the most frequently words
func (ta *TextAnalyzer) prepareTextWords() {
	ta.removeEmptyWords()
	ta.removeSymbols()
	ta.removeEmptyWords()
	ta.sortWords()
}

func (ta *TextAnalyzer) sortWords() {
	SortStrings(ta.Words)
}

func (ta *TextAnalyzer) removeSymbols() {
	regex, _ := regexp.Compile("[\\W]+")
	var words []string
	for _, word := range ta.Words {
		removedSymbols := regex.ReplaceAllString(word, " ")
		splitted := strings.Split(removedSymbols, " ")
		words = append(words, splitted...)
	}
	ta.Words = words
}

func (ta *TextAnalyzer) removeEmptyWords() {
	out := ta.Words[:0]
	for _, word := range ta.Words {
		trimmed := bytes.TrimSpace(word)
		if len(word) != 0 {
			lower := bytes.ToLower(trimmed)
			out = append(out, lower)
		}
	}
	ta.Words = out
}
