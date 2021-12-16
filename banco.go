package main

//Importe de librerias necesarias para el programa
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func banco(nCajeros int, cClientes int, tamCola int, done chan bool) {
	var Clientes int = cClientes
	var colaClientes []string
	var Cajeros []string
	colaClientes = generarColaI(tamCola, &colaClientes)
	Cajeros = generarCajerosI(nCajeros, Cajeros)
	generarClientes(Clientes, tamCola, &colaClientes, &Cajeros, &done)

}

//Corrutina la cual modelara la atencion de clientes mediante los cajeros
func AtenderClientes(colaClientes *[]string, Cajeros *[]string, done *chan bool) {
	colaAux := *colaClientes
	cajerosAux := *Cajeros
	for i := 1; i < len(cajerosAux)+1; {
		var auxNum string = strconv.Itoa(i)
		var auxCaj string = cajerosAux[i-1]
		var Comparador string = "Cajero N°" + auxNum + " - Libre"
		var cliente string
		if auxCaj == Comparador {
			for j := 0; j < len(colaAux); {
				if colaAux[j] != "Libre" {
					cliente = colaAux[j]
					fmt.Print("\nEL Cajero " + auxNum + " Atendera a " + cliente + "\n\n\n")
					cajerosAux[i-1] = "Cajero N°" + auxNum + " - " + cliente
					colaAux[j] = "Libre"
					*Cajeros = cajerosAux
					*colaClientes = colaAux
					imprimirCola(*colaClientes)
					ImprimirCajeros(*Cajeros)
					randomStop()
					fmt.Print("\nEL " + cliente + " Fue Atendido y se marcha del banco \n\n\n")
					cajerosAux[i-1] = "Cajero N°" + auxNum + " - Libre"
					*Cajeros = cajerosAux
					ImprimirCajeros(*Cajeros)
					i++
					j = len(colaAux)
				} else {
					j++
				}
			}
		}
	}
	time.Sleep(2 * time.Second)
	Play(done)

}

//Corrutina la cual simulara el arribo de los clientes al banco
func generarClientes(cClientes int, tamCola int, colaClientes *[]string, Cajeros *[]string, done *chan bool) {
	colaAux := *colaClientes
	for i := 1; i < cClientes+1; {
		colaAux := *colaClientes
		for j := 0; j < len(colaAux); {
			auxC := colaAux[j]
			if j == (len(colaAux)-1) && auxC != "Libre" {
				go AtenderClientes(&colaAux, Cajeros, done)
				Pausa(done)
				j = len(colaAux)
			}
			if auxC == "Libre" {
				var auxNum string = strconv.Itoa(i)
				var aux = ("Cliente N°" + auxNum)
				colaAux[j] = aux
				*colaClientes = colaAux
				fmt.Print("\nLlega el " + aux + "\n\n\n")
				imprimirCola(*colaClientes)
				randomStop()
				i = i + 1
				j = len(colaAux)
			} else {
				j = j + 1
			}
		}

	}
	go AtenderClientes(&colaAux, Cajeros, done)
	Pausa(done)
	fmt.Print("\n''''''''''''Todos los Clientes del Dia Fueron Atendidos''''''''''''''''\n")
}

//randomStop()simulara el tiempo en que se demoran tanto el arribo de los clientes,como el tiempo en que un cajero atiende a un cliente
//los interbalos de tiempo pueden ser editados añadiendo mas argumentos al arreglo r
func randomStop() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := []int{3, 4, 5, 6}
	aux := rand.Intn(len(r))
	aux2 := (r[aux])
	for i := 0; i < aux2; i++ {
		time.Sleep(1 * time.Second)
	}
}

//Carga de datos inciales
func generarColaI(tamCola int, colaClientes *[]string) []string {
	for i := 1; i < tamCola+1; i++ {
		var llenar string = "Libre"
		*colaClientes = append(*colaClientes, llenar)
	}
	imprimirCola(*colaClientes)
	return *colaClientes
}
func generarCajerosI(tamCola int, Cajeros []string) []string {
	for i := 1; i < tamCola+1; i++ {
		var auxNum string = strconv.Itoa(i)
		var auxC = ("Cajero N°" + auxNum + " - Libre")
		Cajeros = append(Cajeros, auxC)
	}
	ImprimirCajeros(Cajeros)
	return Cajeros
}

//Imprimir Por Pantalla
func imprimirCola(colaClientes []string) {
	fmt.Print("                                        COLA DEL BANCO\n")
	fmt.Print("[")
	fmt.Print(colaClientes[0])
	for i := 1; i < len(colaClientes); i++ {
		fmt.Print(", " + colaClientes[i] + "")
	}
	fmt.Print("]  \n")
	fmt.Print("===========================================================================================\n")
}

func ImprimirCajeros(Cajeros []string) {
	fmt.Print("\n                                     	CAJEROS\n")
	fmt.Print("[")
	fmt.Print(Cajeros[0])
	for i := 1; i < len(Cajeros); i++ {
		fmt.Print(", " + Cajeros[i])
	}
	fmt.Print("]  \n")
	fmt.Print("===========================================================================================\n")
}
