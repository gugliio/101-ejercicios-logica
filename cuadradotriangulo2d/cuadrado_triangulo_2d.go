package cuadradotriangulo2d

import "fmt"

/*
 * Crea un programa que dibuje un cuadrado o un triángulo con asteriscos "*".
 * - Indicaremos el tamaño del lado y si la figura a dibujar es una u otra.
 * - EXTRA: ¿Eres capaz de dibujar más figuras?
 */

var (
	Cuadrado  = "Cuadrado"
	Triangulo = "Triangulo"
)

func Execute(size int, polygonType string) {
	if size < 2 {
		fmt.Println("El tamaño debe ser mayor a 1")
		return
	}

	if polygonType == Cuadrado {
		drawSquare(size)
	}

	if polygonType == Triangulo {
		drawTriangle(size)
	}
}

func drawSquare(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func drawTriangle(size int) {
	for i := 1; i <= size; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
