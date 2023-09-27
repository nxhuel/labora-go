package main

import(
	"fmt"
)

type Shape interface {
	Area() float64
}

func main() {
	triangulo := Triangulo {
		Base: 18,
		Altura: 12,
	}
	fmt.Println("Triangulo:")
	Imprimir(triangulo)

	rectangulo := Rectangulo {
		Base: 18,
		Altura: 12,
	}
	fmt.Println("Rectangulo:")
	Imprimir(rectangulo)

	circulo := Circulo {
		Radius: 5,
	}
	fmt.Println("Circulo:")
	Imprimir(circulo)

	trapecio := Trapecio {
		Base: 18,
		Altura: 12,
	}
	fmt.Println("Trapecio:")
	Imprimir(trapecio)

	poligono := Poligono {
		Perimetro: 6,
		Apotema: 8,
	}
	fmt.Println("Poligono:")
	Imprimir(poligono)
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

// Rectangulo

type Rectangulo struct {
	Base float64
	Altura float64
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

// Circulo

type Circulo struct {
	Radius float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Trapecio

type Trapecio struct {
	Base float64
	Altura float64
}

func (t Trapecio) Area() float64 {
	return ((t.Base + t.Base) * t.Altura) / 2
}

// Poligono
type Poligono struct {
	Perimetro float64
	Apotema float64
}

func (p Poligono) Area() float64 {
	return (p.Perimetro * p.Apotema) / 2
}


func Imprimir (s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}