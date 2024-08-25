package sleeptime

import (
	"fmt"
	"time"
)

/*
 * Crea una función que sume 2 números y retorne su resultado pasados
 * unos segundos.
 * - Recibirá por parámetros los 2 números a sumar y los segundos que
 *   debe tardar en finalizar su ejecución.
 * - Si el lenguaje lo soporta, deberá retornar el resultado de forma
 *   asíncrona, es decir, sin detener la ejecución del programa principal.
 *   Se podría ejecutar varias veces al mismo tiempo.
 */

func Execute(a, b int, seconds int) {
	// crear canal
	c := make(chan int)

	// Llamar a la función en una goroutine.
	go executeAsync(3, 5, 2, c)
	go executeAsync(10, 20, 5, c)
	go executeAsync(10, 12, 10, c)

	// Recibir los resultados desde el canal.
	result1 := <-c
	result2 := <-c
	result3 := <-c

	// Mostrar los resultados.
	fmt.Println("Resultado de la primera suma:", result1)
	fmt.Println("Resultado de la segunda suma:", result2)
	fmt.Println("Resultado de la ultima suma:", result3)
}

func executeAsync(a, b int, seconds int, c chan int) {
	time.Sleep(time.Duration(seconds) * time.Second)
	result := a + b
	c <- result
}
