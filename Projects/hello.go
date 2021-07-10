package main

import (
	"fmt"
	//"strconv" //Para convertir int a str
)

func main() {
	// -----DECLARACION DE VARIABLES ----//
	//Palabra reservada para declarar variables var.Para declarar: var nombre_variable tipo_dato
	var x int
	var cadena string
	//var bandera bool
	//var cadenas []string

	x = 23    //Si se declara la variable le asigna valor pero no se usa, no ejecuta el programa.
	y := 4.22 //2da forma de declarar una variable, no necesita declarar el tipo, version tipado dinamico.
	cadena = "sebas"
	cadenas := "Onix"

	fmt.Println("Hello")
	fmt.Println(cadena)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(cadenas)
}
