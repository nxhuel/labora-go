// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type sistemaSolar struct {
// 	nombre string
// 	cantidad int
// }

// func mostrarLunas(planeta sistemaSolar) string {
// 	switch planeta.cantidad {
// 	case 0:
// 		return "No hay lunas"
// 	default:
// 		return fmt.Sprintf("Hay %d lunas", planeta.cantidad)
// 	}
// }

// func main() {
// 	planetas := []sistemaSolar {
// 		{"Mercurio", 0},
// 		{"Venus", 0},
// 		{"Tierra", 1},
// 		{"Marte", 2},
// 		{"Jupiter", 79},
// 		{"Saturno", 83},
// 		{"Urano", 27},
// 		{"Neptuno", 14},
// 	}

// 	sort.Slice(planetas, func (i, j int) bool {
// 		return planetas[i].cantidad < planetas[j].cantidad
// 	})

// 	for i, planeta := range planetas {
// 		fmt.Printf ("%d. %s - %s\n", i+1, planeta.nombre, mostrarLunas(planeta))
// 	}
// }