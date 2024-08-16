package anagrama

import (
	"bytes"
	"sort"
)

func Execute(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	word1Array := []byte(word1)
	sort.Slice(word1Array, func(i, j int) bool {
		return word1Array[i] < word1Array[j]
	})

	word2Array := []byte(word2)
	sort.Slice(word2Array, func(i, j int) bool {
		return word2Array[i] < word2Array[j]
	})

	return bytes.Equal(word1Array, word2Array)
}
