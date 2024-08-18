package revert

/*
 * Crea un programa que invierta el orden de una cadena de texto
 * sin usar funciones propias del lenguaje que lo hagan de forma automática.
 * - Si le pasamos "Hola mundo" nos retornaría "odnum aloH"
 */

func Execute(text string) string {
	var reverse string
	for _, c := range text {
		reverse = string(c) + reverse
	}

	return reverse
}
