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
	words := strings.Split(fileText, " ")

	textAnalyzer := &text.TextAnalyzer{
		Words:                       words,
		MostFrequentlyWordsQuantity: MostFrequentlyWordsQuantity,
	}
	textAnalyzer.FindMostFrequentlyWords()
}
