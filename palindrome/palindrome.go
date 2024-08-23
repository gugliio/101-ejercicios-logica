package palindrome

import "strings"

/*
 * Escribe una función que reciba un texto y retorne verdadero o
 * falso (Boolean) según sean o no palíndromos.
 * Un Palíndromo es una palabra o expresión que es igual si se lee
  * de izquierda a derecha que de derecha a izquierda.
 * NO se tienen en cuenta los espacios, signos de puntuación y tildes.
 * Ejemplo: Ana lleva al oso la avellana.
*/

func Execute(word string) bool {
	var invertedWord string
	for _, char := range word {
		invertedWord = string(char) + invertedWord
	}
	return strings.EqualFold(word, invertedWord)
}
