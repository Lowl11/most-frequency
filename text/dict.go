package text

import "sort"

type Word struct {
	Text  string
	Count int
}

type WordsDictionary struct {
	Words  []string
	Counts []int
}

func (wd *WordsDictionary) AddWord(word string) {
	findWord, findWordIndex := wd.GetWordByText(word)
	if findWord == nil {
		wd.Words = append(wd.Words, word)
		wd.Counts = append(wd.Counts, 1)
	} else {
		wd.Counts[findWordIndex]++
	}
}

func (wd *WordsDictionary) GetWordByText(word string) (*Word, int) {
	wordIndex := sort.SearchStrings(wd.Words, word)
	if wordIndex == len(wd.Words) {
		return nil, 0
	}
	return &Word{
		Text:  wd.Words[wordIndex],
		Count: wd.Counts[wordIndex],
	}, wordIndex
}

func (wd *WordsDictionary) GetWordByIndex(index int) *Word {
	if index < len(wd.Words) {
		return &Word{
			Text:  wd.Words[index],
			Count: wd.Counts[index],
		}
	}
	return nil
}

func (wd *WordsDictionary) Len() int {
	return len(wd.Words)
}

func (wd *WordsDictionary) Less(i, j int) bool {
	return wd.Counts[i] < wd.Counts[j]
}

func (wd *WordsDictionary) Swap(i, j int) {
	wd.Counts[i], wd.Counts[j] = wd.Counts[j], wd.Counts[i]
	wd.Words[i], wd.Words[j] = wd.Words[j], wd.Words[i]
}
