package text

type TextAnalyzer struct {
	Words                       []string
	MostFrequentlyWordsQuantity int
}

// Main method of analyzer which will find the most frequently words
func (ta *TextAnalyzer) FindMostFrequentlyWords() {
	ta.prepareTextWords()
}

func (ta *TextAnalyzer) prepareTextWords() {
	ta.removeEmptyWords()
}

func (ta *TextAnalyzer) removeEmptyWords() {
	out := ta.Words[:0]
	for _, word := range ta.Words {
		if word != "" {
			out = append(out, word)
		}
	}
	ta.Words = out
}
