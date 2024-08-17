package primos_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/primos"
	"github.com/stretchr/testify/require"
)

func Test_Primos(t *testing.T) {
	type args struct {
		number   int
		expected bool
	}
	tt := map[string]args{
		"given 7 return true": {
			number:   7,
			expected: true,
		},
		"given 10 return false": {
			number:   10,
			expected: false,
		},
		"given 1 return false": {
			number:   1,
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := primos.Execute(tc.number)
			require.Equal(t, tc.expected, got)
		})
	}
}
