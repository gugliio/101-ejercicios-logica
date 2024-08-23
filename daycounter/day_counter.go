package daycounter

import (
	"errors"
	"time"
)

/*
 * Crea una función que calcule y retorne cuántos días hay entre dos cadenas
 * de texto que representen fechas.
 * - Una cadena de texto que representa una fecha tiene el formato "dd/MM/yyyy".
 * - La función recibirá dos String y retornará un Int.
 * - La diferencia en días será absoluta (no importa el orden de las fechas).
 * - Si una de las dos cadenas de texto no representa una fecha correcta se
 *   lanzará una excepción.
 */

const _layout = "02/01/2006" // dd/mm/yyyy format for parsing the date strings

var (
	ErrInvalidInitialTime = errors.New("error invalid initial time")
	ErrInvalidFinalTime   = errors.New("error invalid final time")
)

func Execute(initialTime, finalTime string) (int, error) {
	init, err := parseToDate(initialTime)
	if err != nil {
		return 0, ErrInvalidInitialTime
	}

	end, err := parseToDate(finalTime)
	if err != nil {
		return 0, ErrInvalidFinalTime
	}

	diff := end.Sub(init)
	return int(diff.Hours() / 24), nil
}

func parseToDate(date string) (time.Time, error) {
	return time.Parse(_layout, date)
}
