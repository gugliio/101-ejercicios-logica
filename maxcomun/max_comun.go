package maxcomun

import "errors"

/*
 * Crea dos funciones, una que calcule el máximo común divisor (MCD) y otra
 * que calcule el mínimo común múltiplo (mcm) de dos números enteros.
 * - No se pueden utilizar operaciones del lenguaje que
 *   lo resuelvan directamente.
 */

var (
	ErrNegativeNumbers = errors.New("numbers must be positive integers")
)

func ExecuteMDC(num1, num2 int) (int, error) {
	if num1 <= 0 || num2 <= 0 {
		return 0, ErrNegativeNumbers
	}

	for num1 != 0 && num2 != 0 {
		temp := num2
		num2 = num1 % num2
		num1 = temp
	}

	return num1 + num2, nil
}

func ExecuteMCM(num1, num2 int) int {
	val, err := ExecuteMDC(num1, num2)
	if err != nil {
		return 0
	}

	return (num1 * num2) / val
}
