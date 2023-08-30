package main

import "fmt"

func main () {
	var a, b int
	
	fmt.Println("Ingrese el primer numero:")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el segundo numero:")
	fmt.Scanln(&b)

	calcular(&a, &b)

	fmt.Println("Programa finalizado")
}

func calcular (a, b *int) {
	
	suma := *a + *b
	resta := *a - *b
	multiplicacion := *a * *b
	division := *a / *b

	fmt.Println("La suma es:", suma)
	fmt.Println("La resta es:", resta)
	fmt.Println("La multiplicacion es es:", multiplicacion)
	fmt.Println("La division es:", division)
}