package conjuntos_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/conjuntos"
	"github.com/stretchr/testify/require"
)

func Test_Conjuntos(t *testing.T) {
	type args struct {
		elem1      []string
		elem2      []string
		mustSearch bool
		expected   []string
	}
	tt := map[string]args{
		"given elem1 = [1,2,3,4,5] and elem2 = [4,5,6,7,8] and mustSearch = true, return [4,5]": {
			elem1:      []string{"1", "2", "3", "4", "5"},
			elem2:      []string{"4", "5", "6", "7", "8"},
			mustSearch: true,
			expected:   []string{"4", "5"},
		},
		"given elem1 = [1,2,3,4,5] and elem2 = [4,5,6,7,8] and mustSearch = false, return [1,2,3,6,7,8]": {
			elem1:      []string{"1", "2", "3", "4", "5"},
			elem2:      []string{"4", "5", "6", "7", "8"},
			mustSearch: false,
			expected:   []string{"1", "2", "3", "6", "7", "8"},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := conjuntos.Execute(tc.elem1, tc.elem2, tc.mustSearch)
			require.Equal(t, tc.expected, got)
		})
	}
}
