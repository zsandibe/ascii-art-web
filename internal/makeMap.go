package internal

import (
	"strings"
)

func MakeMap(arrByte []byte) map[rune][]string {
	data := strings.ReplaceAll(string(arrByte), "\r", "")
	char := strings.Split(string(data), "\n")
	mapChar := make(map[rune][]string)
	count := 1
	for i := ' '; i <= '~'; i++ {
		mapChar[i] = char[count : count+8]
		count = count + 9
	}
	return mapChar
}
