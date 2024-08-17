package main

import (
	"fmt"

	"github.com/gugliio/101-ejercicios-logica/anagrama"
	"github.com/gugliio/101-ejercicios-logica/fibonacci"
	"github.com/gugliio/101-ejercicios-logica/fizzbuzz"
	"github.com/gugliio/101-ejercicios-logica/poligonos"
	"github.com/gugliio/101-ejercicios-logica/primos"
)

func main() {
	fmt.Println("Reto 01 - Fizzbuzz")
	fizzbuzz.Execute()

	fmt.Println("Reto 02 - Anagrama")
	fmt.Println(anagrama.Execute("amor", "roma"))

	fmt.Println("Reto 03 - Fibonacci")
	fibonacci.Execute()

	fmt.Println("Reto 04 - Numeros Primos")
	primos.Execute(3)

	fmt.Println("Reto 05 - Poligono")
	fmt.Println(poligonos.Execute(string(poligonos.TipoCuadrado), 2, 2))
}
