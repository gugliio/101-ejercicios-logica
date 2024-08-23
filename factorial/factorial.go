package factorial

/*
 * Escribe una función que calcule y retorne el factorial de un número dado
 * de forma recursiva.
 */

func Execute(number int) int {
	if number <= 0 {
		return 1
	}

	result := 1
	for i := number; i > 0; i-- {
		result = result * i
	}

	return result
}
