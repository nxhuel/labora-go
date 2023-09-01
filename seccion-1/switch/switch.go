// package main

// import (
// 	"fmt"
// )

// func main() {
// 	dias := [7]string {"Domingo", "Lunes", "Martes", "Miecoles", "Jueves", "Viernes", "Sabado"}

// 	for i, dia := range dias {
// 		fmt.Printf("[%d] = %s \n", i, dia)
// 	}

// 	var diaElegido int
// 	fmt.Printf("Elija un dia de la semana: ")
// 	fmt.Scan(&diaElegido)

// 	switch diaElegido {
// 	case 1:
// 		fmt.Println("Lunes")
// 	case 2:
// 		fmt.Println("Martes")
// 	case 3:
// 		fmt.Println("Miercoles")
// 	case 4:
// 		fmt.Println("Jueves")
// 	case 5:
// 		fmt.Println("Viernes")
// 	case 6:
// 		fmt.Println("Sabado")
// 	case 0:
// 		fmt.Println("Domingo")
// 	default:
// 		fmt.Println("El valor ingresado no es un día válido.")
// 	}
// }
