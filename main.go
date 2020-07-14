package main

import (
	"io/ioutil"
	"most-frequency/text"
	"strings"
)

// Quantity of the most frequently words
const MostFrequentlyWordsQuantity = 20

func main() {
	file, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err.Error())
	}

	fileText := string(file)
	words := splitWords(fileText)

	textAnalyzer := &text.TextAnalyzer{
		Words:                       words,
		MostFrequentlyWordsQuantity: MostFrequentlyWordsQuantity,
	}
	textAnalyzer.FindMostFrequentlyWords()
}

func splitWords(text string) []string {
	var words []string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		lineWords := strings.Split(line, " ")
		words = append(words, lineWords...)
	}
	return words
}
