package text

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type TextAnalyzer struct {
	Words                       []string
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

	mostFrequentWord := make([]Word, 0, ta.MostFrequentlyWordsQuantity)
	sort.Sort(dict)
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
	ta.sortWords()
}

func (ta *TextAnalyzer) sortWords() {
	sort.Strings(ta.Words)
}

func (ta *TextAnalyzer) removeEmptyWords() {
	out := ta.Words[:0]
	regex, _ := regexp.Compile("[\\W]+")
	for _, word := range ta.Words {
		if word != "" {
			preparedWord := strings.ToLower(strings.TrimSpace(word))
			preparedWord = regex.ReplaceAllString(preparedWord, "")
			out = append(out, preparedWord)
		}
	}
	ta.Words = out
}
