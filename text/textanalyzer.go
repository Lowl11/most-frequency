package text

import (
	"fmt"
	"regexp"
	"strings"
)

type TextAnalyzer struct {
	Words                       []string
	MostFrequentlyWordsQuantity int
}

// Main method of analyzer which will find the most frequently words
func (ta *TextAnalyzer) FindMostFrequentlyWords() {
	ta.prepareTextWords()
	for _, word := range ta.Words {
		fmt.Println(word)
	}
}

func (ta *TextAnalyzer) prepareTextWords() {
	ta.removeEmptyWords()
	ta.removeSymbols()
}

func (ta *TextAnalyzer) removeSymbols() {
	regex, _ := regexp.Compile("[^a-zA-Z]+")
	for i := 0; i < len(ta.Words); i++ {
		ta.Words[i] = regex.ReplaceAllString(ta.Words[i], "")
	}
}

func (ta *TextAnalyzer) removeEmptyWords() {
	out := ta.Words[:0]
	for _, word := range ta.Words {
		if word != "" {
			out = append(out, strings.TrimSpace(word))
		}
	}
	ta.Words = out
}
