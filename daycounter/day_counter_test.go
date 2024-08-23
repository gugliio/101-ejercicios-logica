package daycounter_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/daycounter"
	"github.com/stretchr/testify/require"
)

func Test_Day_Counter(t *testing.T) {
	type args struct {
		initDate     string
		finalDate    string
		expectedDays int
	}
	tt := map[string]args{
		"given initial date 01/01/1990 and 01/02/1990 should return 31 days and no errors": {
			initDate:     "01/01/1990",
			finalDate:    "01/02/1990",
			expectedDays: 31,
		},
		"given initial date 07/08/2024 and 06/08/2026 should return 729 days and no errors": {
			initDate:     "07/08/2024",
			finalDate:    "06/08/2026",
			expectedDays: 729,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := daycounter.Execute(tc.initDate, tc.finalDate)
			require.Equal(t, tc.expectedDays, got)
			require.NoError(t, err)
		})
	}
}

func Test_Day_Counter_Error(t *testing.T) {
	type args struct {
		initDate    string
		finalDate   string
		expectedErr error
	}
	tt := map[string]args{
		"given initial date 01-01-1990 and 01/02/1990 should return error invalid initial date": {
			initDate:    "01-01-1990",
			finalDate:   "01/02/1990",
			expectedErr: daycounter.ErrInvalidInitialTime,
		},
		"given initial date 01/01/1990 and 01-02-1990 should return error invalid end date": {
			initDate:    "01/01/1990",
			finalDate:   "01-02-1990",
			expectedErr: daycounter.ErrInvalidFinalTime,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			_, err := daycounter.Execute(tc.initDate, tc.finalDate)
			require.Error(t, err)
			require.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
