package exercise

import (
	"strings"
)

func WordCount(str string) map[string]int {
	count := make(map[string]int)

	for _, word := range strings.Split(str, " ") {
		count[word]++
	}

	return count
}
