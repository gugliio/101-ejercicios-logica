package timeconverter_test

import (
	"testing"
	"time"

	"github.com/gugliio/101-ejercicios-logica/timeconverter"
	"github.com/stretchr/testify/require"
)

func Test_Time_Converter(t *testing.T) {
	type args struct {
		days, hour, minute, second int
		expected                   time.Duration
	}
	tt := map[string]args{
		"given 1 day should return 86400000 milliseconds": {
			days:     1,
			hour:     0,
			minute:   0,
			second:   0,
			expected: time.Duration(86400000) * time.Millisecond,
		},
		"given 1 hour should return 3600000 milliseconds": {
			days:     0,
			hour:     1,
			minute:   0,
			second:   0,
			expected: time.Duration(3600000) * time.Millisecond,
		},
		"given 1 minute should return 60000 milliseconds": {
			days:     0,
			hour:     0,
			minute:   1,
			second:   0,
			expected: time.Duration(60000) * time.Millisecond,
		},
		"given 1 second should return 1000 milliseconds": {
			days:     0,
			hour:     0,
			minute:   0,
			second:   1,
			expected: time.Duration(1000) * time.Millisecond,
		},
		"given 1 day, 1 hour, 1 minute, 1 second should return 90061000 milliseconds": {
			days:     1,
			hour:     1,
			minute:   1,
			second:   1,
			expected: time.Duration(90061000) * time.Millisecond,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := timeconverter.Execute(tc.days, tc.hour, tc.minute, tc.second)
			require.Equal(t, got, tc.expected)
		})
	}

}
