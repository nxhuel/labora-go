//  Sistema que procesa múltiples mensajes en paralelo.

package main

import (
	"fmt"
	"time"
)

type Mensaje struct {
	email string
	sms string
	notificacionPush string
	contenido string
}

func ingresoDeMensajes(mensajes []Mensaje, c chan<- Mensaje) {
	for _, mensaje := range mensajes {
		c <- mensaje
	}
	close(c)
}

func mensajesVisibles(c <-chan Mensaje) {
	for mensaje := range c {
		fmt.Println("Email:", mensaje.email)
		fmt.Println("SMS:", mensaje.sms)
		fmt.Println("Notificación Push:", mensaje.notificacionPush)
		fmt.Println("Contenido:", mensaje.contenido)
	}
}

func main() {
	c := make(chan Mensaje)
	
	var emailUsuario string
	fmt.Print("Ingrese su email: ")
	fmt.Scan(&emailUsuario)

	mensaje := Mensaje {
		email: emailUsuario,
		sms: "SMS anuncio",
		notificacionPush: "Estas son notificacions push!",
		contenido: "Este es un contenido automatico",
	}

	mensajes := []Mensaje{mensaje}

	go ingresoDeMensajes(mensajes, c)
	go mensajesVisibles(c)

	time.Sleep(time.Second)
}