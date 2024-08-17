package main

import (
	"fmt"

	"github.com/gugliio/101-ejercicios-logica/anagrama"
	"github.com/gugliio/101-ejercicios-logica/fibonacci"
	"github.com/gugliio/101-ejercicios-logica/fizzbuzz"
)

func main() {
	fmt.Println("Reto 01 - Fizzbuzz")
	fizzbuzz.Execute()

	fmt.Println("Reto 02 - Anagrama")
	fmt.Println(anagrama.Execute("amor", "roma"))

	fmt.Println("Reto 03 - Fibonacci")
	fibonacci.Execute()
}
