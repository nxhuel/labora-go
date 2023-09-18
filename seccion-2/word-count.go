package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I ate a donut. Then I ate another donut."
	resultado := contarPalabras(s)
	fmt.Println(resultado)
}

func contarPalabras(s string) map[string]int {
	palabras := strings.Fields(s)
	resultado := make(map[string]int)

	for _, palabra := range palabras {
		resultado[palabra]++
	}
	return resultado
}