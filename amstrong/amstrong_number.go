package amstrong

import (
	"fmt"
	"math"
	"strconv"
)

/*
 * Escribe una función que calcule si un número dado es un número de Armstrong
 * (o también llamado narcisista).
 * Si no conoces qué es un número de Armstrong, debes buscar información
 * al respecto.
 */

func Execute(number int) bool {
	stringNumber := fmt.Sprint(number)
	quantityOfChars := float64(len(stringNumber))
	var result int

	for _, char := range stringNumber {
		floatChar, _ := strconv.ParseFloat(string(char), 64)
		result = result + int(math.Pow(float64(floatChar), quantityOfChars))
	}

	return number == result
}
