package primos

import "fmt"

/*
 * Escribe un programa que se encargue de comprobar si un número es o no primo.
 * Hecho esto, imprime los números primos entre 1 y 100.
 */

func Execute(number int) bool {
	primo := isPrimo(number)

	for i := 0; i <= 100; i++ {
		if isPrimo(i) {
			fmt.Println(i)
		}
	}

	return primo
}

func isPrimo(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
