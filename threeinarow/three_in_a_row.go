package threeinarow

/*
 * Crea una función que analice una matriz 3x3 compuesta por "X" y "O"
 * y retorne lo siguiente:
 * - "X" si han ganado las "X"
 * - "O" si han ganado los "O"
 * - "Empate" si ha habido un empate
 * - "Nulo" si la proporción de "X", de "O", o de la matriz no es correcta.
 *   O si han ganado los 2.
 * Nota: La matriz puede no estar totalmente cubierta.
 * Se podría representar con un vacío "", por ejemplo.
 */

func Execute(matrix [][]string) string {
	isDiagonal, value := isDiagonal(matrix)
	if isDiagonal {
		return value
	}

	isLine, value := isHorizontalLine(matrix)
	if isLine {
		return value
	}

	isVertical, value := isVerticalLine(matrix)
	if isVertical {
		return value
	}

	if isFull(matrix) {
		return "Empate"
	}

	return "Nulo"
}

func isFull(matrix [][]string) bool {
	for _, row := range matrix {
		for _, cell := range row {
			if cell != "X" && cell != "O" {
				return false
			}
		}
	}
	return true
}

func isDiagonal(matrix [][]string) (bool, string) {
	if matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] && matrix[0][0] != " " {
		return true, matrix[0][0]
	}
	if matrix[0][2] == matrix[1][1] && matrix[1][1] == matrix[2][0] && matrix[0][2] != " " {
		return true, matrix[0][2]
	}
	return false, ""
}

func isHorizontalLine(matrix [][]string) (bool, string) {
	for _, row := range matrix {
		if row[0] == row[1] && row[1] == row[2] && row[0] != " " {
			return true, row[0]
		}
	}
	return false, ""
}

func isVerticalLine(matrix [][]string) (bool, string) {
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == matrix[1][i] && matrix[1][i] == matrix[2][i] && matrix[0][i] != " " {
			return true, matrix[0][i]
		}
	}
	return false, ""
}
