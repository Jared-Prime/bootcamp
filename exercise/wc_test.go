package exercise

import (
	"testing"
)

var (
	text          = "a b a a c d e c a"
	expectedCount = map[string]int{"a": 4, "b": 1, "c": 2, "d": 1}
)

func TestWordCount(t *testing.T) {
	result := WordCount(text)
	for word, count := range expectedCount {
		if res := result[word]; res != count {
			t.Errorf("expected %d count for %s, got only %d", count, word, res)
		}
	}
}
