package morsecode_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/morsecode"

	"github.com/stretchr/testify/require"
)

func Test_CodeMorse(t *testing.T) {
	type args struct {
		input    string
		expected string
	}
	tt := map[string]args{
		"given the word perro return .--. . .-. .-. ---": {
			input:    "perro",
			expected: ".--. . .-. .-. ---",
		},
		"given the word gato return -... .-.-. .- -..": {
			input:    "gato",
			expected: "--. .- - ---",
		},
		"given the code --. .- - --- should return gato": {
			input:    "--. .- - --- ",
			expected: "gato",
		},
		"given the code .--. . .-. .-. ---  should return perro": {
			input:    ".--. . .-. .-. --- ",
			expected: "perro",
		},
		"given the code -.. --- ... .--. . .-. .-. --- ...  should return dos perros": {
			input:    "-.. --- ... / .--. . .-. .-. --- ...",
			expected: "dos perros",
		},
		"given the words dos perros should return the code -.. --- ...  .--. . .-. .-. --- ...": {
			input:    "dos perros",
			expected: "-.. --- ... / .--. . .-. .-. --- ...",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := morsecode.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
