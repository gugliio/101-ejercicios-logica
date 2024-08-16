package anagrama_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/anagrama"
	"github.com/stretchr/testify/require"
)

func Test_Anagrama(t *testing.T) {
	type args struct {
		word1    string
		word2    string
		expected bool
	}
	tt := map[string]args{
		"given roma and amor should return true": {
			word1:    "roma",
			word2:    "amor",
			expected: true,
		},
		"given hello and olleh should return true": {
			word1:    "hello",
			word2:    "olleh",
			expected: true,
		},
		"given hello and world should return false": {
			word1:    "hello",
			word2:    "world",
			expected: false,
		},
		"given perros and world should return false": {
			word1:    "perros",
			word2:    "world",
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := anagrama.Execute(tc.word1, tc.word2)
			require.Equal(t, tc.expected, got)
		})
	}
}
