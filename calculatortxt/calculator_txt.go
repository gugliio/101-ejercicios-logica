package calculatortxt

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
 * Lee el fichero "Challenge21.txt" incluido en el proyecto, calcula su
 * resultado e imprímelo.
 * - El .txt se corresponde con las entradas de una calculadora.
 * - Cada línea tendrá un número o una operación representada por un
 *   símbolo (alternando ambos).
 * - Soporta números enteros y decimales.
 * - Soporta las operaciones suma "+", resta "-", multiplicación "*"
 *   y división "/".
 * - El resultado se muestra al finalizar la lectura de la última
 *   línea (si el .txt es correcto).
 * - Si el formato del .txt no es correcto, se indicará que no se han
 *   podido resolver las operaciones.
 */

var (
	ErrReadingFile      = errors.New("err reading file")
	ErrDivisionByZero   = errors.New("err division by zero")
	ErrInvalidOperation = errors.New("err invalid operation")
	ErrParsingFloat     = errors.New("err parsing float")

	digitCheck = regexp.MustCompile(`^[0-9]+$`)
)

func Execute(filePath string) (float64, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return 0, ErrReadingFile
	}
	lines := strings.Split(string(bytes), "\n")

	var result float64
	var action string
	for _, value := range lines {
		if digitCheck.MatchString(value) {
			number, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return 0, ErrParsingFloat
			}
			result, err = updateResult(result, number, action)
			if err != nil {
				return 0, err
			}
		}
		action = value
	}
	return result, nil
}

func updateResult(result, number float64, action string) (float64, error) {
	if result == 0 {
		return number, nil
	}
	return calculate(result, number, action)
}

func calculate(result, number float64, action string) (float64, error) {
	if action != "" {
		switch action {
		case "+":
			result += number
		case "-":
			result -= number
		case "*":
			result *= number
		case "/":
			if number == 0 {
				return 0, ErrDivisionByZero
			}
			result /= number
		default:
			return 0, ErrInvalidOperation
		}
	}
	return result, nil
}
