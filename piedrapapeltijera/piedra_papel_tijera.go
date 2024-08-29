package piedrapapeltijera

import "errors"

/*
 * Crea un programa que calcule quien gana más partidas al piedra,
 * papel, tijera.
 * - El resultado puede ser: "Player 1", "Player 2", "Tie" (empate)
 * - La función recibe un listado que contiene pares, representando cada jugada.
 * - El par puede contener combinaciones de "R" (piedra), "P" (papel)
 *   o "S" (tijera).
 * - Ejemplo. Entrada: [("R","S"), ("S","R"), ("P","S")]. Resultado: "Player 2".
 */

var (
	ErrInvalidInput = errors.New("error: invalid input")
	Ganador1        = "Ganador 1"
	Ganador2        = "Ganador 2"
	Empate          = "Empate"
)

func Execute(jug1 string, jug2 string) (string, error) {
	if jug1 != "piedra" && jug1 != "papel" && jug1 != "tijera" {
		return "", ErrInvalidInput
	}

	if jug2 != "piedra" && jug2 != "papel" && jug2 != "tijera" {
		return "", ErrInvalidInput
	}

	if jug1 == jug2 {
		return Empate, nil
	}

	switch jug1 {
	case "piedra":
		if jug2 == "tijera" {
			return Ganador1, nil
		}
		return Ganador2, nil
	case "papel":
		if jug2 == "piedra" {
			return Ganador1, nil
		}
		return Ganador2, nil
	case "tijera":
		if jug2 == "papel" {
			return Ganador1, nil
		}
		return Ganador2, nil
	}
	return "", nil // This should never be reached, but added for completeness
}
