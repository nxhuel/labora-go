package main

import (
	"fmt"
	"sort"
)

type Estudiantes struct {
	nombre string
	nota int
	codigo string
}

func crearEstudiantes() []Estudiantes {
	estudiante := []Estudiantes {}
	var e Estudiantes

	var numDeEstudiantes int
	fmt.Print("Ingresa la cantidad de estudiantes: ")
	fmt.Scan(&numDeEstudiantes)
	fmt.Printf("El total de estudiantes son %d\n", numDeEstudiantes)

	for i := 1; i <= numDeEstudiantes; i++ {
		fmt.Print("\nIngresa el nombre del alumno: ")
		fmt.Scan(&e.nombre)
	
		fmt.Printf("Ingresa la nota del alumno %s: ", e.nombre)
		fmt.Scan(&e.nota)
	
		fmt.Printf("Ingresa el codigo del alumno %s: ", e.nombre)
		fmt.Scan(&e.codigo)

		estudiante = append(estudiante, e)
	}

	for index, value := range estudiante{
		fmt.Printf("\nEstudiante %d\n", index+1)
        fmt.Printf("Nombre: %s\n", value.nombre)
        fmt.Printf("Nota: %d\n", value.nota)
        fmt.Printf("Codigo: %s\n", value.codigo)
	}

	return estudiante
}

func ordenarLista(estudiante []Estudiantes)  {
	var ordenado int
	fmt.Println("Elija el orden que desee:\n1. Nombre\n2. Nota\n3. Codigo")
	fmt.Scan(&ordenado)

	if ordenado == 1 {
		sort.Slice(estudiante, func(i, j int) bool {
			return estudiante[i].nombre < estudiante[j].nombre
		})		
		fmt.Println("Lista ordenada por nombres")
		for _, nombresOrdenados := range estudiante {
			fmt.Printf("Nombre: %s, Nota: %d, Codigo: %s\n", nombresOrdenados.nombre, nombresOrdenados.nota, nombresOrdenados.codigo)
		}
	} else if ordenado == 2 {
		sort.Slice(estudiante, func(i, j int) bool {
			return estudiante[i].nota < estudiante[j].nota
		})
		fmt.Println("Lista ordenada por notas")
		for _, notasOrdenados := range estudiante {
			fmt.Printf("Nombre: %s, Nota: %d, Codigo: %s\n", notasOrdenados.nombre, notasOrdenados.nota, notasOrdenados.codigo)
		}
	} else if ordenado == 3 {
		sort.Slice(estudiante, func(i, j int) bool {
			return estudiante[i].codigo < estudiante[j].codigo
		})
		fmt.Println("Lista ordenada por codigos")
		for _, codigosOrdenados := range estudiante {
			fmt.Printf("Nombre: %s, Nota: %d, Codigo: %s\n", codigosOrdenados.nombre, codigosOrdenados.nota, codigosOrdenados.codigo)
		}
	} else {
		fmt.Println(nil)
	}
}

func buscarEstudiante(estudiante []Estudiantes) {
	var aprobado int
	fmt.Print("Ingrese (1) si desea ver la lista de aprobados o ingrese (2) si desea ver la lista de desaprobados: \n")
	fmt.Scan(&aprobado)

	if aprobado == 1 {
		for i := 0; i < len(estudiante); i++ {
			if estudiante[i].nota >= 4 {
				fmt.Printf("%s aprobado\n", estudiante[i].nombre)
			}
		}
	} else if aprobado == 2 {
		for i := 0; i < len(estudiante); i++ {
			if estudiante[i].nota < 4 {
				fmt.Printf("%s desaprobado\n", estudiante[i].nombre)
			}
		}
	} else {
		fmt.Println(nil)
	}
}

func main() {
	for {
		var centilenta string
		fmt.Print("Ingrese '*' para salir: ")
		fmt.Scan(&centilenta)
		if centilenta == "*" {
			break
		} else {
		estudiantes := crearEstudiantes()
		ordenarLista(estudiantes)
		buscarEstudiante(estudiantes)
		}
	}
}

