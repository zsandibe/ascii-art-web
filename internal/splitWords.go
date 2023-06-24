package internal

import (
	"errors"
	"strings"
)

func SplitWords(data string) ([]string, error) {
	var Error error
	for _, char := range data {
		if !(char >= 0 && char <= '~') {
			Error = errors.New("Not ascii chars")
		}
	}
	check := 0
	sep := []string{"\\n", "\n"}
	words := strings.Split(strings.ReplaceAll(data, "\r", ""), sep[0])
	var newWords []string
	for _, word := range words {
		subWords := strings.Split(word, sep[1])
		newWords = append(newWords, subWords...)
	}
	words = newWords

	for _, word := range words {
		if len(word) == 0 {
			check++
		}
	}
	if len(words) == check {
		words = words[:len(words)-1]
	}
	return words, Error
}
