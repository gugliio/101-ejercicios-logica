package poligonos_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/poligonos"
	"github.com/stretchr/testify/require"
)

func Test_Poligono(t *testing.T) {
	type args struct {
		tipoPoligono poligonos.Type
		base         float64
		altura       float64
		expected     float64
	}
	tt := map[string]args{
		"given a polygon type triangle with a base 3 and altura 1 expected 1.5": {
			base:         3,
			altura:       1,
			expected:     1.5,
			tipoPoligono: poligonos.TipoTriangulo,
		},
		"given a polygon type cuadrado with a base 2 and altura 2 expected 4": {
			base:         2,
			altura:       2,
			expected:     4,
			tipoPoligono: poligonos.TipoCuadrado,
		},
		"given a polygon type rectangulo with a base 6 and altura 2 expected 12": {
			base:         6,
			altura:       2,
			expected:     12,
			tipoPoligono: poligonos.TipoRectangulo,
		},
		"given a polygon type invalid return 0": {
			base:         6,
			altura:       2,
			expected:     0,
			tipoPoligono: poligonos.Type("invalid"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := poligonos.Execute(string(tc.tipoPoligono), tc.base, tc.altura)
			require.Equal(t, tc.expected, got)
		})
	}
}
