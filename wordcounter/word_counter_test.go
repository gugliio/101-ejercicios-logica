package wordcounter_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/wordcounter"
	"github.com/stretchr/testify/require"
)

func Test_Word_Counter(t *testing.T) {
	type args struct {
		input    string
		word     string
		expected int
	}
	tt := map[string]args{
		"given input on upper case and twice the word musica should return 2": {
			input:    "Estaba escuchando MUSICA en la casa de la musica",
			word:     "musica",
			expected: 2,
		},
		"given input on lower case and special char twice the word musica should return 2": {
			input:    "Estaba escuchando musica en la casa de la música",
			word:     "musica",
			expected: 2,
		},
		"given input on lower case and special char should return 1": {
			input:    "Hello world, test",
			word:     "world",
			expected: 1,
		},
		"given input on upper case and the word with camel case should return 1": {
			input:    "Hello World",
			word:     "World",
			expected: 1,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := wordcounter.Execute(tc.input, tc.word)
			require.Equal(t, tc.expected, got)
		})
	}
}
