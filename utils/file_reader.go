package utils

import (
	"bufio"
	"most-frequency/text"
	"os"
)

func ReadFile(fileName string) text.String {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	scanner := bufio.NewScanner(file)
	channel := make(chan text.String)
	defer close(channel)

	go func(words chan text.String) {
		for scanner.Scan() {
			words <- text.String(scanner.Text())
		}
	}(channel)

	var fileText text.String
	for word := range channel {
		fileText = text.AppendString(fileText, word)
	}

	return fileText
}
