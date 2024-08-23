package amstrong_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/amstrong"
	"github.com/stretchr/testify/require"
)

func Test_Amstrong(t *testing.T) {
	type args struct {
		input    int
		expected bool
	}
	tt := map[string]args{
		"given 2 should return true": {
			input:    2,
			expected: true,
		},
		"given 10 should return false": {
			input:    10,
			expected: false,
		},
		"given 30 should return false": {
			input:    30,
			expected: false,
		},
		"given 153 should return true": {
			input:    153,
			expected: true,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := amstrong.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
