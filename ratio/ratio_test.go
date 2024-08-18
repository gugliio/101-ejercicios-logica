package ratio_test

import (
	"testing"

	"github.com/gugliio/101-ejercicios-logica/ratio"
	"github.com/stretchr/testify/require"
)

func Test_Ratio(t *testing.T) {
	url := "https://raw.githubusercontent.com/mouredev/mouredev/master/mouredev_github_profile.png"
	expected := "5:2"
	got := ratio.Execute(url)
	require.Equal(t, expected, got)
}
