package text

import (
	"bytes"
	"sort"
)

type String []byte

func AppendString(current String, new ...String) String {
	newLength := len(current) + sumOfLengths(new...)
	result := make([]byte, newLength, newLength)
	copy(result[:], current)
	previousIndex := len(current)
	for _, value := range new {
		copy(result[previousIndex:], value)
		previousIndex += len(value)
	}
	return result
}

func sumOfLengths(strings ...String) int {
	sum := 0
	for _, value := range strings {
		sum += len(value)
	}
	return sum
}

func ConvertBytesToString(array [][]byte) []String {
	var result []String
	for _, bytes := range array {
		result = append(result, String(bytes))
	}
	return result
}

func ConcatenateArrays(array1 []String, array2 []String) []String {
	newLength := len(array1) + len(array2)
	result := make([]String, newLength, newLength)
	copy(result[:], array1[:])
	copy(result[len(array1):], array2)
	return result
}

func SortStrings(strings Strings) {
	sort.Sort(&strings)
}

type Strings []String

func (ss *Strings) Len() int {
	return len(*ss)
}

func (ss *Strings) Less(i, j int) bool {
	comparing := bytes.Compare((*ss)[i], (*ss)[j])
	if comparing == 1 {
		return true
	}
	return false
}

func (ss *Strings) Swap(i, j int) {
	(*ss)[i], (*ss)[j] = (*ss)[j], (*ss)[i]
}
