// package main

// import "fmt"

// func main () {

// 	fmt.Println("1. Restaurante 1")
// 	fmt.Println("2. Restaurante 2")
// 	fmt.Println("3. Restaurante 3")

// 	var restaurante int
// 	fmt.Println("Ingrese el numero del restaurante: ")
// 	fmt.Scanln(&restaurante)

// 	var calificacion float32
// 	fmt.Println("Ingrese la calificacion que desea darle al restaurante: ")
// 	fmt.Scanln(&calificacion)

// 	fmt.Printf("El restaurante %d tiene una calificación de %.1f.\n", restaurante, calificacion)

// 	-----------------------

// 	for i := 1; i < 4; i++ {
// 		if i == 1 {
// 			fmt.Printf("El resturante %d tiene una calificacion de 4.5.\n", i)
// 		} else if i == 2 {
// 			fmt.Printf("El resturante %d tiene una calificacion de 6.5.\n", i)
// 		} else if i == 3 {
// 			fmt.Printf("El resturante %d tiene una calificacion de 8.5.\n", i)
// 		}
// 	}
// }