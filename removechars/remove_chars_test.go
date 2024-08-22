package removechars_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/removechars"
	"github.com/stretchr/testify/require"
)

func Test_Compare_strings(t *testing.T) {
	type args struct {
		input1    string
		input2    string
		expected1 string
		expected2 string
	}
	tt := map[string]args{
		"given input1 perro and input 2 gato should return out1 perr and gat": {
			input1:    "perro",
			input2:    "gato",
			expected1: "perr",
			expected2: "gat",
		},
		"given input1 copa and input 2 padre should return out1 co and dre": {
			input1:    "copa",
			input2:    "padre",
			expected1: "co",
			expected2: "dre",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got1, got2 := removechars.Execute(tc.input1, tc.input2)
			require.Equal(t, tc.expected1, got1)
			require.Equal(t, tc.expected2, got2)
		})
	}
}
