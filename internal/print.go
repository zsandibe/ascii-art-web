package internal

func Print(words []string, mapChar map[rune][]string) string {
	var result string
	for i := 0; i <= len(words)-1; i++ {
		if len(words[i]) == 0 {
			result = result + "\n"
			continue
		}
		for j := 0; j < 8; j++ {
			for _, v := range words[i] {
				result = result + mapChar[v][j]
			}
			result = result + "\n"

		}
	}
	return result
}
