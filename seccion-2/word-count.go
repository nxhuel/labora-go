// V.1

// package main

// import(
// 	"fmt"
// 	"strings"
// )

// func contador(){
// 	var s string
// 	s = "I ate a donut. Then I ate another donut."
// 	palabras := strings.Fields(s)

// 	resultado := make(map[string]int)

// 	for _, palabra := range palabras {
// 		resultado[palabra]++
// 	}
// 	fmt.Println(resultado)
// }

// func main(){
// 	contador()
// }

// V.2

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solicitud() string{
	var oracion string
	fmt.Println("Ingrese una oracion:")
	lectura := bufio.NewReader(os.Stdin)
	oracion, _ = lectura.ReadString('\n')
	return oracion
}

func contador(){
	var x string
	x = solicitud()
	palabras := strings.Fields(x)

	resultado := make(map[string]int)

	for _, palabra := range palabras {
		resultado[palabra]++
	}
	fmt.Println(resultado)
}

func main(){
	contador()
}