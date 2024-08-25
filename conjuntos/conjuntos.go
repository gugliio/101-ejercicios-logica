package conjuntos

import "slices"

/*
 * Crea una función que reciba dos array, un booleano y retorne un array.
 * - Si el booleano es verdadero buscará y retornará los elementos comunes
 *   de los dos array.
 * - Si el booleano es falso buscará y retornará los elementos no comunes
 *   de los dos array.
 * - No se pueden utilizar operaciones del lenguaje que
 *   lo resuelvan directamente.
 */

func Execute(elem1, elem2 []string, mustSearch bool) []string {
	var res []string
	for _, s := range elem1 {
		if mustSearch && slices.Contains(elem2, s) {
			res = append(res, s)
		}

		if !mustSearch && !slices.Contains(elem2, s) {
			res = append(res, s)
		}
	}

	for _, s := range elem2 {
		if !mustSearch && !slices.Contains(elem1, s) {
			res = append(res, s)
		}
	}

	return res
}
