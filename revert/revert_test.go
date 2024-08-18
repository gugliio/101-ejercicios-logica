package revert_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/revert"
	"github.com/stretchr/testify/require"
)

func Test_Revert(t *testing.T) {
	type args struct {
		input    string
		expected string
	}
	tt := map[string]args{
		"given Hola mundo should return odnum aloH": {
			input:    "Hola mundo",
			expected: "odnum aloH",
		},
		"given hello should return olleh": {
			input:    "hello",
			expected: "olleh",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := revert.Execute(tc.input)
			require.Equal(t, tc.expected, got)
		})
	}
}
