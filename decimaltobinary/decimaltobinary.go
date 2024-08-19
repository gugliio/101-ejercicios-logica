package decimaltobinary

import (
	"fmt"
	"strconv"
)

/*
 * Crea un programa se encargue de transformar un número
 * decimal a binario sin utilizar funciones propias del lenguaje que lo hagan directamente.
 */

func Execute(number int) string {
	var binary string

	if number == 0 {
		return "0"
	}

	for number > 0 {
		binary = fmt.Sprintf("%s%s", strconv.Itoa(number%2), binary)
		number /= 2
	}

	return binary
}
