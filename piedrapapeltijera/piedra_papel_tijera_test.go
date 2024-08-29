package piedrapapeltijera_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/piedrapapeltijera"
	"github.com/stretchr/testify/require"
)

func Test_Piedra_Papel_Tijera(t *testing.T) {
	type args struct {
		jug1     string
		jug2     string
		expected string
	}
	tt := map[string]args{
		"if jug1 plays tijera and jug2 plays papel win jug 1": {
			jug1:     "tijera",
			jug2:     "papel",
			expected: piedrapapeltijera.Ganador1,
		},
		"if jug2 plays tijera and jug1 plays papel win jug 2": {
			jug1:     "papel",
			jug2:     "tijera",
			expected: piedrapapeltijera.Ganador2,
		},
		"if jug1 plays piedra and jug2 plays papel win jug 2": {
			jug1:     "piedra",
			jug2:     "papel",
			expected: piedrapapeltijera.Ganador2,
		},
		"if jug2 plays piedra and jug1 plays papel win jug 1": {
			jug1:     "papel",
			jug2:     "piedra",
			expected: piedrapapeltijera.Ganador1,
		},
		"if jug1 plays piedra and jug2 plays tijera win jug 1": {
			jug1:     "piedra",
			jug2:     "tijera",
			expected: piedrapapeltijera.Ganador1,
		},
		"if jug2 plays piedra and jug1 plays tijera win jug 2": {
			jug1:     "tijera",
			jug2:     "piedra",
			expected: piedrapapeltijera.Ganador2,
		},
		"if both players play piedra return empate": {
			jug1:     "piedra",
			jug2:     "piedra",
			expected: piedrapapeltijera.Empate,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := piedrapapeltijera.Execute(tc.jug1, tc.jug2)
			require.Nil(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}

func Test_Piedra_Papel_Tijera_Error(t *testing.T) {
	type args struct {
		jug1        string
		jug2        string
		expectedErr error
	}
	tt := map[string]args{
		"if jug1 plays invalid return error": {
			jug1:        "invalid",
			jug2:        "papel",
			expectedErr: piedrapapeltijera.ErrInvalidInput,
		},
		"if jug2 plays invalid return error": {
			jug1:        "papel",
			jug2:        "invalid",
			expectedErr: piedrapapeltijera.ErrInvalidInput,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			_, err := piedrapapeltijera.Execute(tc.jug1, tc.jug2)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}
