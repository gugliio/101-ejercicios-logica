package threeinarow_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/threeinarow"
	"github.com/stretchr/testify/require"
)

func Test_Three_in_a_Row(t *testing.T) {
	type args struct {
		inputMatrix [][]string
		expected    string
	}
	tt := map[string]args{
		"Test with a winning diagonal row - X": {
			inputMatrix: [][]string{{"X", "O", "X"}, {"O", "X", " "}, {" ", " ", "X"}},
			expected:    "X",
		},
		"Test with a winning diagonal row - O": {
			inputMatrix: [][]string{{"O", " ", " "}, {" ", "O", " "}, {" ", " ", "O"}},
			expected:    "O",
		},
		"Test with a winning horizontal row": {
			inputMatrix: [][]string{{"X", "X", "X"}, {" ", " ", " "}, {" ", " ", " "}},
			expected:    "X",
		},
		"Test with a winning horizontal row - O": {
			inputMatrix: [][]string{{"O", "O", "O"}, {" ", " ", " "}, {" ", " ", " "}},
			expected:    "O",
		},
		"Test with a winning vertical row": {
			inputMatrix: [][]string{{"X", " ", " "}, {"X", "X", " "}, {"X", " ", " "}},
			expected:    "X",
		},
		"Test with a winning vertical row - O": {
			inputMatrix: [][]string{{"O", " ", " "}, {"O", "O", " "}, {"O", " ", " "}},
			expected:    "O",
		},
		"Test with a winning anti-diagonal row": {
			inputMatrix: [][]string{{" ", " ", "X"}, {" ", "X", " "}, {"X", " ", " "}},
			expected:    "X",
		},
		"Test with a winning anti-diagonal row - O": {
			inputMatrix: [][]string{{" ", " ", "O"}, {" ", "O", " "}, {"O", " ", " "}},
			expected:    "O",
		},
		"Test with a winning row with mixed elements": {
			inputMatrix: [][]string{{"X", "O", "O"}, {"O", "X", "X"}, {"X", "O", "X"}},
			expected:    "X",
		},
		"Test with a winning row with empty spaces": {
			inputMatrix: [][]string{{"X", " ", " "}, {" ", " ", " "}, {" ", " ", " "}},
			expected:    "Nulo",
		},
		"Test with a draw": {
			inputMatrix: [][]string{{"X", "O", "X"}, {"O", "O", "X"}, {"O", "X", "O"}},
			expected:    "Empate",
		},
		"Test with a game not ended": {
			inputMatrix: [][]string{{"X", " ", " "}, {"O", " ", " "}, {" ", " ", " "}},
			expected:    "Nulo",
		},
	}
	for _, tc := range tt {
		got := threeinarow.Execute(tc.inputMatrix)
		require.Equal(t, tc.expected, got)
	}
}
