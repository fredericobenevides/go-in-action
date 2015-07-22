package words

import "strings"

// Calcula
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
