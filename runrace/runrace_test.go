package runrace_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/runrace"
	"github.com/stretchr/testify/require"
)

func Test_Run_Race(t *testing.T) {
	type args struct {
		jumpRun  []string
		pista    []string
		expected bool
	}
	tt := map[string]args{
		"given run and jump return true": {
			jumpRun:  []string{"run", "jump"},
			pista:    []string{"_", "|"},
			expected: true,
		},
		"given run and run return false": {
			jumpRun:  []string{"run", "run"},
			pista:    []string{"_", "|"},
			expected: false,
		},
		"given jump and jump return false": {
			jumpRun:  []string{"jump", "jump"},
			pista:    []string{"_", "|"},
			expected: false,
		},
		"given run, jump, run, run, jump return true": {
			jumpRun:  []string{"run", "jump", "run", "run", "jump"},
			pista:    []string{"_", "|", "_", "_", "|"},
			expected: true,
		},
	}
	for _, tc := range tt {
		got := runrace.Execute(tc.jumpRun, tc.pista)
		require.Equal(t, tc.expected, got)
	}

}
