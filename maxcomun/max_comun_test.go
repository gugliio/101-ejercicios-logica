package maxcomun_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/maxcomun"
	"github.com/stretchr/testify/require"
)

func Test_ExecuteMDC(t *testing.T) {
	type args struct {
		input1      int
		input2      int
		expected    int
		expectedErr error
	}
	tt := map[string]args{
		"given 10 and 20 should return 10": {
			input1:   10,
			input2:   20,
			expected: 10,
		},
		"given 0 and 10 should return error invalid number": {
			input1:      0,
			input2:      10,
			expected:    0,
			expectedErr: maxcomun.ErrNegativeNumbers,
		},
		"given 20 and 0 should return error invalid number": {
			input1:      20,
			input2:      0,
			expected:    0,
			expectedErr: maxcomun.ErrNegativeNumbers,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := maxcomun.ExecuteMDC(tc.input1, tc.input2)
			require.Equal(t, tc.expected, got)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func Test_ExecuteMDM(t *testing.T) {
	type args struct {
		input1   int
		input2   int
		expected int
	}
	tt := map[string]args{
		"given 10 and 20 should return 20": {
			input1:   10,
			input2:   20,
			expected: 20,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := maxcomun.ExecuteMCM(tc.input1, tc.input2)
			require.Equal(t, tc.expected, got)
		})
	}
}
