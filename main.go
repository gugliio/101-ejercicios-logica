package main

import (
	"fmt"

	"github.com/gugliio/101-ejercicios-logica/anagrama"
	"github.com/gugliio/101-ejercicios-logica/decimaltobinary"
	"github.com/gugliio/101-ejercicios-logica/fibonacci"
	"github.com/gugliio/101-ejercicios-logica/fizzbuzz"
	"github.com/gugliio/101-ejercicios-logica/poligonos"
	"github.com/gugliio/101-ejercicios-logica/primos"
	"github.com/gugliio/101-ejercicios-logica/ratio"
	"github.com/gugliio/101-ejercicios-logica/revert"
	"github.com/gugliio/101-ejercicios-logica/wordcounter"
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

	fmt.Println("Reto 06 - Ratio")
	url := "https://raw.githubusercontent.com/mouredevmouredev/master/mouredev_github_profile.png"
	fmt.Println(ratio.Execute(url))

	fmt.Println("Reto 07 - Reverse string")
	fmt.Println(revert.Execute("Hello"))

	fmt.Println("Reto 08 - Word counter")
	fmt.Println(wordcounter.Execute("Hello World", "World"))

	fmt.Println("Reto 09 - Decimal to binary")
	fmt.Println(decimaltobinary.Execute(10))
}
