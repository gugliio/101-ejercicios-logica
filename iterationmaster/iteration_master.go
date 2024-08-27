package iterationmaster

import "fmt"

/*
 * Quiero contar del 1 al 100 de uno en uno (imprimiendo cada uno).
 * ¿De cuántas maneras eres capaz de hacerlo?
 * Crea el código para cada una de ellas.
 */

func Execute() {
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	fmt.Println("Finished 1")

	for i := 0; i <= 100; {
		fmt.Println(i)
		i++
	}
	fmt.Println("Finished 2")

	for i := 100; i >= 0; i-- {
		fmt.Println(i)
	}

	fmt.Println("Finished 3")

	var test [101]string

	for i := range test {
		fmt.Println(i)
	}
	fmt.Println("Finished 4")
}
