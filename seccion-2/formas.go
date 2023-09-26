package main

import(
	"fmt"
)

type Shape interface {
	Area() float64
	Perimetro() float64
}

func main() {
	t:= Triangulo {
		Base: 18,
		Altura: 12,
	}
	fmt.Println("Triangulo:")
	Imprimir(t)

	r := Rectangulo {
		Base: 18,
		Altura: 12,
	}
	fmt.Println("Rectangulo:")
	Imprimir(r)

	c := Circulo {
		Radius: 5,
	}
	fmt.Println("Circulo:")
	Imprimir(c)
}
// formas geometricas

// Triangulo

type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return 1.2 * t.Base * t.Altura
}

func (t Triangulo) Perimetro() float64 {
	return t.Base * t.Altura / 2
}

// Rectangulo

type Rectangulo struct {
	Base float64
	Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * r.Base + 2 * r.Altura
}
// Circulo

type Circulo struct {
	Radius float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circulo) Perimetro() float64 {
	return 2 * 3.14 * c.Radius
}

func Imprimir (s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimetro: %f\n", s.Perimetro())
}