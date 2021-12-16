package main

import (
	"fmt"
)

func main() {

	menu := "Bienvenido\nDesea revisar:\n [1] Certamen 2.A y 2.B \n [2] Certamen 2 Extendido\n"
	var i int
	fmt.Print(menu)
	fmt.Scanln(&i)
	switch i {
	case 1:
		fmt.Print("Comiezo Programa de Prueba Corrutinas\n")
		done := make(chan bool, 1)
		go CorrutinaA(done)
	case 2:
		fmt.Print("Comiezo Programa de Simulador De Banco\n")
		done1 := make(chan bool, 1)
		go banco(4, 16, 4, done1)
	}
	var wait string
	fmt.Scanln(&wait)
}
