package main

import (
	"fmt"

	"github.com/gugliio/101-ejercicios-logica/amstrong"
	"github.com/gugliio/101-ejercicios-logica/anagrama"
	"github.com/gugliio/101-ejercicios-logica/calculatortxt"
	"github.com/gugliio/101-ejercicios-logica/daycounter"
	"github.com/gugliio/101-ejercicios-logica/decimaltobinary"
	"github.com/gugliio/101-ejercicios-logica/expequilibradas"
	"github.com/gugliio/101-ejercicios-logica/factorial"
	"github.com/gugliio/101-ejercicios-logica/fibonacci"
	"github.com/gugliio/101-ejercicios-logica/fizzbuzz"
	"github.com/gugliio/101-ejercicios-logica/morsecode"
	"github.com/gugliio/101-ejercicios-logica/palindrome"
	"github.com/gugliio/101-ejercicios-logica/poligonos"
	"github.com/gugliio/101-ejercicios-logica/primos"
	"github.com/gugliio/101-ejercicios-logica/ratio"
	"github.com/gugliio/101-ejercicios-logica/removechars"
	"github.com/gugliio/101-ejercicios-logica/revert"
	"github.com/gugliio/101-ejercicios-logica/runrace"
	"github.com/gugliio/101-ejercicios-logica/sleeptime"
	"github.com/gugliio/101-ejercicios-logica/threeinarow"
	"github.com/gugliio/101-ejercicios-logica/timeconverter"
	"github.com/gugliio/101-ejercicios-logica/uppercase"
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

	fmt.Println("Reto 10 - Codigo Morse")
	fmt.Println(morsecode.Execute("Ariel Ortega"))

	fmt.Println("*Reto 11 - Expresiones equilibradas")
	fmt.Println(expequilibradas.Execute("{}"))

	fmt.Println("Reto 12 - Remove chars")
	fmt.Println(removechars.Execute("perro", "gatos"))

	fmt.Println("Reto 13 - Palindrome")
	fmt.Println(palindrome.Execute("perro"))

	fmt.Println("Reto 14 - Factorial")
	fmt.Println(factorial.Execute(10))

	fmt.Println("Reto 15 - Amstrong Number")
	fmt.Println(amstrong.Execute(50))

	fmt.Println("Reto 16 - Days Counter")
	fmt.Println(daycounter.Execute("25/05/1990", "09/12/2018"))

	fmt.Println("Reto 17 - Uppercase")
	fmt.Println(uppercase.Execute("hello world"))

	fmt.Println("Reto 18 - Run Race")
	fmt.Println(runrace.Execute([]string{"run", "jump"}, []string{"_", "|"}))

	fmt.Println("Reto 19 - Three in a Row")
	fmt.Println(threeinarow.Execute([][]string{{"X", "O", "X"}, {"O", "X", " "}, {" ", " ", "X"}}))

	fmt.Println("Reto 20 - Time converter")
	fmt.Println(timeconverter.Execute(1, 2, 3, 4))

	fmt.Println("Reto 21 - Sleep time")
	sleeptime.Execute(2, 1, 2)

	fmt.Println("Reto 22 - Calculator txt")
	fmt.Println(calculatortxt.Execute("./calculatortxt/calculator.txt"))
}
