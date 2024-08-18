package wordcounter

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var ruleRemoveAccents = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

func Execute(input, word string) int {
	wordCounter := 0

	inputAsSlice := strings.Split(input, " ")

	for _, inputWord := range inputAsSlice {
		inputWord = removeAccents(inputWord)
		inputWord = removeChars(inputWord)
		if strings.ToLower(inputWord) == word {
			wordCounter++
		}
	}
	return wordCounter
}

func removeAccents(word string) string {
	res, _, _ := transform.String(ruleRemoveAccents, word)
	return res
}

func removeChars(word string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(word, "")
}
