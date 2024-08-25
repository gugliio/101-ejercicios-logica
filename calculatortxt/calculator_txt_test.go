package calculatortxt_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/calculatortxt"
	"github.com/stretchr/testify/require"
)

func Test_Calculator(t *testing.T) {
	type args struct {
		filePath         string
		expectedResponse float64
		expectedErr      error
	}
	tt := map[string]args{
		"given the right path return ": {
			filePath:         "./calculator.txt",
			expectedResponse: 18.5,
		},
		"given an invalid file path return error": {
			filePath:         "invalid",
			expectedResponse: 0,
			expectedErr:      calculatortxt.ErrReadingFile,
		},
		"given a division by zero return error division by zero": {
			filePath:         "./div_by_zero.txt",
			expectedResponse: 0,
			expectedErr:      calculatortxt.ErrDivisionByZero,
		},
		"given a random char return error invalid operation": {
			filePath:         "./invalid_operator.txt",
			expectedResponse: 0,
			expectedErr:      calculatortxt.ErrInvalidOperation,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := calculatortxt.Execute(tc.filePath)
			require.Equal(t, tc.expectedResponse, got)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
