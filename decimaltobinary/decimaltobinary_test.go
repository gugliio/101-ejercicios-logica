package decimaltobinary_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/decimaltobinary"
	"github.com/stretchr/testify/require"
)

func Test_DecimalToBinary(t *testing.T) {
	type args struct {
		decimalInput int
		expected     string
	}
	tt := map[string]args{
		"given 10 as input should return 1010": {
			decimalInput: 10,
			expected:     "1010",
		},
		"given 75 as input should return 1001011": {
			decimalInput: 75,
			expected:     "1001011",
		},
		"given 0 as input should return 0": {
			decimalInput: 0,
			expected:     "0",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := decimaltobinary.Execute(tc.decimalInput)
			require.Equal(t, tc.expected, got)
		})
	}
}
