package removechars

import "strings"

/*
 * Crea una función que reciba dos cadenas como parámetro (str1, str2)
 * e imprima otras dos cadenas como salida (out1, out2).
 * - out1 contendrá todos los caracteres presentes en la str1 pero NO
 *   estén presentes en str2.
 * - out2 contendrá todos los caracteres presentes en la str2 pero NO
 *   estén presentes en str1.
 */

func Execute(str1, str2 string) (string, string) {
	var out1, out2 string
	for _, s := range str1 {
		if !strings.ContainsRune(str2, s) {
			out1 += string(s)
		}
	}

	for _, v := range str2 {
		if !strings.ContainsRune(str1, v) {
			out2 += string(v)
		}
	}

	return out1, out2
}
