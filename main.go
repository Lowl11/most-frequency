package main

import (
	"bytes"
	"most-frequency/text"
	"most-frequency/utils"
)

// Quantity of the most frequently words
const MostFrequentlyWordsQuantity = 20

func main() {
	fileName := "text.txt"
	var fileText text.String = utils.ReadFile(fileName)
	words := splitWords(fileText)

	textAnalyzer := &text.TextAnalyzer{
		Words:                       words,
		MostFrequentlyWordsQuantity: MostFrequentlyWordsQuantity,
	}
	textAnalyzer.FindMostFrequentlyWords()
}

func splitWords(fileText text.String) []text.String {
	var words []text.String
	space := []byte(" ")
	lines := bytes.Split(fileText, space)
	for _, line := range lines {
		lineWords := bytes.Split(line, space)
		stringWords := text.ConvertBytesToString(lineWords)
		words = text.ConcatenateArrays(words, stringWords)
	}
	return words
}
