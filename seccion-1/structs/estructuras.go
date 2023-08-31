package main

import (
	"fmt"
)

func main() {
	personaUno := persona {
		nombre: "Juan",
		edad: 30,
		ciudad: "Madrid",
		telefono: 5551234,
	}

	personaDos := persona {
		nombre: "Ana",
		edad: 25,
		ciudad: "Barcelona",
		telefono: 5551234,
	}

	fmt.Println("Persona 1:", personaUno)
	fmt.Println("Persona 2:", personaDos)

	cambiarCiudar(&personaUno, "Valencia")

	fmt.Println("Persona 1 con ciudad actualizada:", personaUno)
}

type persona struct {
	nombre string
	edad int
	ciudad string
	telefono int
}

func cambiarCiudar(p *persona, ciudad string) {
	if p.ciudad != ciudad {
		p.ciudad = ciudad
		fmt.Println("Ciudad cambiada a", ciudad)
	} else {
		fmt.Println("La persona ya vive en", ciudad)
	}
}

