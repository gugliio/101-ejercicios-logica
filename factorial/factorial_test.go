package factorial_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/factorial"
	"github.com/stretchr/testify/require"
)

func Test_Factorial(t *testing.T) {
	type args struct {
		input    int
		expected int
	}
	tt := map[string]args{
		"given 10 as input should return 3628800": {
			input:    10,
			expected: 3628800,
		},
		"given 5 as input should return 120": {
			input:    5,
			expected: 120,
		},
		"given 0 as input should return 0": {
			input:    0,
			expected: 1,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := factorial.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
