package uppercase

import (
	"strings"
)

/*
 * Crea una función que reciba un String de cualquier tipo y se encargue de
 * poner en mayúscula la primera letra de cada palabra.
 * - No se pueden utilizar operaciones del lenguaje que
 *   lo resuelvan directamente.
 */

func Execute(input string) string {
	sliceWords := strings.Split(input, " ")
	result := ""
	for _, word := range sliceWords {
		word = strings.ToUpper(word[:1]) + word[1:]
		result = updateResponse(result, word)
	}
	return result
}

func updateResponse(response, word string) string {
	if response == "" {
		return word
	}
	return response + " " + word
}
