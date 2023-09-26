// package main

// import(
// 	"fmt"
// 	"strings"
// )

// func mapa() {
// 	btc := 50

// 	usuarios := []string {
// 		"Matthew",
// 		"Sarah", 
// 		"Augustus", 
// 		"Heidi", 
// 		"Emilie",
// 		"Peter", 
// 		"Giana", 
// 		"Adriano", 
// 		"Aaron", 
// 		"Elizabeth",
// 	}

// 	distribucion := make(map[string]int, len(usuarios))

// 	for _, usuario := range usuarios {
// 		vocales := 0
// 		for _, letra := range strings.ToLower(usuario) {
// 			switch letra {
// 			case 'a':
// 				vocales += 1
// 			case 'e':
// 				vocales += 1
// 			case 'i':
// 				vocales += 2
// 			case 'o':
// 				vocales += 3
// 			case 'u':
// 				vocales += 4
// 			}
// 		}
// 		if vocales > 0 {
// 			if vocales > 10 {
// 				vocales = 10
// 			}
// 			if btc >= vocales {
// 				distribucion[usuario] = vocales
// 				btc = btc - vocales
// 			} else {
// 				distribucion[usuario] = btc
// 				btc = 0
// 			}
// 		}
// 	}

// 	fmt.Println(distribucion)

// 	fmt.Printf("Monedas restantes: %d\n", btc)

// }

// func main() {
// 	mapa()
// }