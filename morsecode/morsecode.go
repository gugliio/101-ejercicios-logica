package morsecode

import (
	"strings"
)

var charToMorse = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	" ": "/",
}

func Execute(word string) string {
	if IsCodeMorse(word) {
		return morseToLetters(word)
	}

	return letterToMorse(word)
}

func IsCodeMorse(word string) bool {
	return strings.Contains(word, ".") || strings.Contains(word, "-")
}

func letterToMorse(word string) string {
	var morse string
	for _, char := range word {
		letter := strings.ToUpper(string(char))
		if morseCode, ok := charToMorse[letter]; ok {
			morse += morseCode + " "
		}
	}
	return strings.TrimSpace(morse)
}

func morseToLetters(word string) string {
	var words string
	signs := strings.Split(word, " ")
	for _, sign := range signs {
		for k, v := range charToMorse {
			if v == sign {
				words += strings.ToLower(k)
			}
		}
	}
	return strings.TrimSpace(words)
}
