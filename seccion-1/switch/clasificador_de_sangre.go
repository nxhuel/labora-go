// Escribe un programa en Go que clasifique a las personas según su tipo de sangre. 
// Crea una función que tome como entrada el tipo de sangre como cadena y devuelva una cadena con la descripción de la clasificación. 
// Utiliza la declaración switch para determinar la clasificación según el tipo de sangre ingresado. Por ejemplo, si se ingresa "AB+", el programa debe devolver "Grupo sanguíneo AB, factor Rh positivo".

package main

import(
	"fmt"
)

func clasificarSangre(tipoSangre string) string {
	switch tipoSangre {
	case "A+":
		return "Grupo sanguíneo A, factor Rh positivo."
	case "B+":
		return "Grupo sanguíneo B, factor Rh positivo."
	case "AB+":
		return "Grupo sanguíneo AB, factor Rh positivo."
	case "O+":
		return "Grupo sanguíneo O, factor Rh positivo."
	case "A-":
		return "Grupo sanguíneo A, factor Rh negativo."
	case "B-":
		return "Grupo sanguíneo B, factor Rh negativo."
	case "AB-":
		return "Grupo sanguíneo AB, factor Rh negativo."
	case "O-":
		return "Grupo sanguíneo O, factor Rh negativo."
	default:
		return "Tipo de sangre desconocido."
	}
}

func main() {
	fmt.Println("Este programa clasifica a las personas segun su tipo de sangre.")

	var tipos string
	fmt.Println("Elija su tipo de sangre:")
	fmt.Scan(&tipos)

	fmt.Printf("Su tipo de sangre es %s, %s", tipos, clasificarSangre(tipos))
}