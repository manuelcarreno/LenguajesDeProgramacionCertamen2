package main

import (
	"fmt"
	"time"
)

func CorrutinaA(done chan bool) {
	for i := 1; i < 11; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		if i == 3 {
			go CorrutinaB(done)
			Pausa(&done)
		}
		if i == 7 {
			go CorrutinaC(done)
			Pausa(&done)
		}
	}
	fmt.Println("FIN de la Corrutina A")
}

func CorrutinaB(done chan bool) {
	fmt.Println("Pausa de Corrutina A\nLLamada a Corrutina B")
	letras := [4]string{"a", "b", "c", "d"}
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(letras[i])
	}
	fmt.Println("Termino de Corrutina B\nPlay a Corrutina A")
	Play(&done)
}

func CorrutinaC(done chan bool) {
	fmt.Println("Pausa de Corrutina A\nLLamada a Corrutina C")
	colores := [4]string{"azul", "rojo", "amarillo", "verde"}
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(colores[i])
	}
	fmt.Println("Termino de Corrutina C\nPlay a Corrutina A")
	Play(&done)
}

func Pausa(done *chan bool) {
	<-*done
}
func Play(done *chan bool) {
	*done <- true
}

