package uppercase_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/uppercase"
	"github.com/stretchr/testify/require"
)

func Test_Uppercase_Letters(t *testing.T) {
	type args struct {
		input    string
		expected string
	}
	tt := map[string]args{
		"given the word hello return HellO WorlD": {
			input:    "hello world",
			expected: "Hello World",
		},
		"given the cancion de hielo y fuego return CancioN DE HielO Y FuegO": {
			input:    "cancion de hielo y fuego",
			expected: "Cancion De Hielo Y Fuego",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := uppercase.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
