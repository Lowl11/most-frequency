package main

import (
	"io/ioutil"
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

	textAnalyzer := &TextAnalyzer{
		Words:                       strings.Split(fileText, " "),
		MostFrequentlyWordsQuantity: MostFrequentlyWordsQuantity,
	}
	textAnalyzer.FindMostFrequentlyWords()
}
