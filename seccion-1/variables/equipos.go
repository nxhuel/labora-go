package main

import "fmt"

func main () {

	fmt.Println("1. Equipo Uno")
	fmt.Println("2. Equipo Dos")
	fmt.Println("3. Equipo Tres")
	fmt.Println("4. Equipo Cuatro")
	fmt.Println("5. Equipo Cinco")

	var equipos int
	fmt.Println("Elija su equipo: ")
	fmt.Scan(&equipos)
	
	if equipos == 1 {
		fmt.Println("Su equipo ha ganado 28 titulos")
	} else if equipos == 2 {
		fmt.Println("Su equipo ha ganado 8 titulos")
	} else if equipos == 3 {
		fmt.Println("Su equipo ha ganado 18 titulos")
	} else if equipos == 4 {
		fmt.Println("Su equipo ha ganado 4 titulos")
	} else if equipos == 5 {
		fmt.Println("Su equipo ha ganado 0 titulos")
	} 
}