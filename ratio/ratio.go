package ratio

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

/*
 * Crea un programa que se encargue de calcular el aspect ratio de una
 * imagen a partir de una url.
 * - Url de ejemplo:
 *   https://raw.githubusercontent.com/mouredevmouredev/master/mouredev_github_profile.png
 * - Por ratio hacemos referencia por ejemplo a los "16:9" de una
 *   imagen de 1920*1080px.
 */

func Execute(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	divisor := gcdRatio(img.Width, img.Height)
	return fmt.Sprintf("%d:%d", img.Width/divisor, img.Height/divisor)
}

func gcdRatio(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRatio(b, a%b)
}
