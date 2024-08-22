package expequilibradas_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/expequilibradas"
	"github.com/stretchr/testify/require"
)

func Test_ExpresionesEquilibradas(t *testing.T) {
	type args struct {
		input    string
		expected bool
	}
	tt := map[string]args{
		"given the expresion } should return false": {
			input:    "}",
			expected: false,
		},
		"given the expresion {} should return false": {
			input:    "{}",
			expected: true,
		},
		"given the expresion {(} should return false": {
			input:    "{(}",
			expected: false,
		},
		"given the expresion {()} should return false": {
			input:    "{()}",
			expected: true,
		},
		"given the expresion { [ a * ( c + d ) ] - 5 } should return true": {
			input:    "{ [ a * ( c + d ) ] - 5 }",
			expected: true,
		},
		"given the expresion { a * ( c + d ) ] - 5 } should return false": {
			input:    "{ a * ( c + d ) ] - 5 }",
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := expequilibradas.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
