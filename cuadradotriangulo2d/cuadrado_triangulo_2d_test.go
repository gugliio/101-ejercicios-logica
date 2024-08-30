package cuadradotriangulo2d_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/cuadradotriangulo2d"
)

func Test_Cuadrado_Triangulo_2D(t *testing.T) {
	type args struct {
		size        int
		polygonType string
	}
	tt := map[string]args{
		"given 10 as input should print square": {
			size:        10,
			polygonType: cuadradotriangulo2d.Cuadrado,
		},
		"given 10 as input should print triangle": {
			size:        10,
			polygonType: cuadradotriangulo2d.Triangulo,
		},
		"given 0 as input should print warn message": {
			size:        1,
			polygonType: cuadradotriangulo2d.Triangulo,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			cuadradotriangulo2d.Execute(tc.size, tc.polygonType)
		})
	}
}
