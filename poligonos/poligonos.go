package poligonos

import "errors"

/*
 * Crea una única función (importante que sólo sea una) que sea capaz
 * de calcular y retornar el área de un polígono.
 * - La función recibirá por parámetro sólo UN polígono a la vez.
 * - Los polígonos soportados serán Triángulo, Cuadrado y Rectángulo.
 * - Imprime el cálculo del área de un polígono de cada tipo.
 */
type Type string

const (
	TipoTriangulo  Type = "Triangulo"
	TipoCuadrado   Type = "Cuadrado"
	TipoRectangulo Type = "Rectangulo"
)

var ErrInvalidType = errors.New("err el tipo de polígono no es válido")

type Poligono interface {
	area() float64
}

type triangulo struct {
	altura, base float64
}

func (t triangulo) area() float64 {
	area := (t.base * t.altura) / 2
	return area
}

func NewTriangulo(altura, base float64) Poligono {
	return triangulo{
		altura: altura,
		base:   base,
	}
}

type cuadrado struct {
	altura, base float64
}

func (t cuadrado) area() float64 {
	area := t.base * t.altura
	return area
}

func NewCuadrado(altura, base float64) Poligono {
	return cuadrado{
		altura: altura,
		base:   base,
	}
}

type rectangulo struct {
	altura, base float64
}

func (t rectangulo) area() float64 {
	area := t.base * t.altura
	return area
}

func NewRectangulo(altura, base float64) Poligono {
	return rectangulo{
		altura: altura,
		base:   base,
	}
}

func Execute(tipoPoligono string, base, altura float64) float64 {
	switch tipoPoligono {
	case string(TipoTriangulo):
		triangulo := NewTriangulo(altura, base)
		return getArea(triangulo)
	case string(TipoCuadrado):
		cuadrado := NewCuadrado(altura, base)
		return getArea(cuadrado)
	case string(TipoRectangulo):
		rectangulo := NewRectangulo(altura, base)
		return getArea(rectangulo)
	default:
		return 0
	}
}

func getArea(poligon Poligono) float64 {
	return poligon.area()
}
