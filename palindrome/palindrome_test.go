package palindrome_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/palindrome"
	"github.com/stretchr/testify/require"
)

func Test_Palindrome(t *testing.T) {
	type args struct {
		input    string
		expected bool
	}
	tt := map[string]args{
		"given input Pedro return false": {
			input:    "Pedro",
			expected: false,
		},
		"given input Neuquen return true": {
			input:    "Neuquen",
			expected: true,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := palindrome.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
