package main

import (
	"fmt"
	"time"
)

func GorrutinaA(done chan bool) {
	for i := 1; i < 11; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		if i == 3 {
			go GorrutinaB(done)
			Pausa(&done)
		}
		if i == 7 {
			go GorrutinaC(done)
			Pausa(&done)
		}
	}
	fmt.Println("FIN de la Corrutina A")
}

func GorrutinaB(done chan bool) {
	fmt.Println("Pausa de Gorrutina A\nLLamada a Gorrutina B")
	letras := [4]string{"a", "b", "c", "d"}
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(letras[i])
	}
	fmt.Println("Termino de Gorrutina B\nPlay a Gorrutina A")
	Play(&done)
}

func GorrutinaC(done chan bool) {
	fmt.Println("Pausa de Gorrutina A\nLLamada a Gorrutina C")
	colores := [4]string{"azul", "rojo", "amarillo", "verde"}
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(colores[i])
	}
	fmt.Println("Termino de Gorrutina C\nPlay a Gorrutina A")
	Play(&done)
}

func Pausa(done *chan bool) {
	<-*done
}
func Play(done *chan bool) {
	*done <- true
}
