package main

import (
	"fmt"
	"sort"
	"math"
)

type Persona struct {
	nombre string
	edad int
	altura int
	peso int
}

func ordenarPersonas() []Persona {

	personas := []Persona{}

	for i := 0; i < 5; i++ {
		var p Persona

		fmt.Printf("\nIngrese los datos de la persona:\n")

		fmt.Print("Nombre: ")
		fmt.Scan(&p.nombre)
	
		fmt.Print("Edad: ")
		fmt.Scan(&p.edad)
	
		fmt.Print("Altura: ")
		fmt.Scan(&p.altura)
	
		fmt.Print("Peso: ")
		fmt.Scan(&p.peso)

		// fmt.Printf("Su nombre es: %s, su edad es: %d, su altura es: %d y su peso es: %dkg", p.nombre, p.edad, p.altura, p.peso)

		personas = append(personas, p)
	}

	fmt.Println("Lista de personas desordenadas")
	for _, p :=  range personas {
		fmt.Printf("Nombre: %s, Edad: %d, Altura: %d, Peso: %dkg\n", p.nombre, p.edad, p.altura, p.peso)
	}

	var n int
	fmt.Println("Elija el orden que desee:\n1. Nombre\n2. Edad\n3. Altura\n4. Peso")
	fmt.Scan(&n)

	if n == 2 {
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].edad < personas[j].edad
		})
	
		fmt.Println("Lista ordenada por edades")
		for _, persona := range personas {
			fmt.Printf("%s, %d, %d, %d\n", persona.nombre, persona.edad, persona.altura, persona.peso)
		}
	} else if n == 3 {
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].altura < personas[j].altura
		})
	
		fmt.Println("Lista ordenada por altura")
		for _, persona := range personas {
			fmt.Printf("%s, %d, %d, %d\n", persona.nombre, persona.edad, persona.altura, persona.peso)
		}
	} else if n == 4 {
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].peso < personas[j].peso
		})
	
		fmt.Println("Lista ordenada por peso")
		for _, persona := range personas {
			fmt.Printf("%s, %d, %d, %d\n", persona.nombre, persona.edad, persona.altura, persona.peso)
		}
	} else if n == 1 {
		sort.Slice(personas, func(i, j int) bool { 
			return personas[i].nombre < personas[j].nombre 
		})
			
		fmt.Println("Lista ordenada por nombre")
		for _, persona := range personas {
			fmt.Printf("%s, %d, %d, %d\n", persona.nombre, persona.edad, persona.altura, persona.peso)
		}
	} else {
		fmt.Println(nil)
	}

	return personas
}

func buscarPersona(personas []Persona) {
	var nombre string
	fmt.Println("\nIngrese el nombre de la persona que desea buscar: ")
	fmt.Scan(&nombre)

	for _, busqueda := range personas {
		if busqueda.nombre == nombre {
			fmt.Printf("\n%s, %d, %d, %d\n", busqueda.nombre, busqueda.edad, busqueda.altura, busqueda.peso)
			return
		} else {
			fmt.Println(nil)
			return
		}
	}
}

func calculoIMC(personas []Persona) {

	// var peso float64
	// peso = 72
	// var altura float64
	// altura = 172

	for _, masa := range personas {
		imc := float64(masa.peso) / math.Pow(float64(masa.altura)/100, 2)

		fmt.Printf("\nEl indice de masa corporal de %s es: %f", masa.nombre, imc)

        if imc < 18.5 {
            fmt.Println("\nBajo peso")
        } else if imc >= 18.5 && imc <= 24.9 {
            fmt.Println("\nPeso normal")
        } else if imc >= 25 && imc <= 29.9 {
            fmt.Println("\nSobrepeso")
        } else if imc > 30{
            fmt.Println("\nObesidad")
        }
	}
	return
}

func main() {
	fmt.Printf("Este programa permite el ingreso de los datos de 5 personas y calcula su IMC.")
	
	personas := ordenarPersonas()
	buscarPersona(personas)
	calculoIMC(personas)
}