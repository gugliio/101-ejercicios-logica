package expequilibradas

/*
 * Crea un programa que comprueba si los paréntesis, llaves y corchetes
 * de una expresión están equilibrados.
 * - Equilibrado significa que estos delimitadores se abren y cieran
 *   en orden y de forma correcta.
 * - Paréntesis, llaves y corchetes son igual de prioritarios.
 *   No hay uno más importante que otro.
 * - Expresión balanceada: { [ a * ( c + d ) ] - 5 }
 * - Expresión no balanceada: { a * ( c + d ) ] - 5 }
 */

var symbols = map[rune]rune{'{': '}', '[': ']', '(': ')'}

func Execute(expression string) bool {
	stack := []rune{}
	for _, char := range expression {
		if _, ok := symbols[char]; ok {
			stack = append(stack, char)
		} else if isClosing(char) {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if closing, exists := symbols[last]; !exists || char != closing {
				return false
			}
		}
	}

	return len(stack) == 0
}

func isClosing(s rune) bool {
	return s == '}' || s == ']' || s == ')'
}
