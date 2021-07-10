package main

import (
	"fmt"
)

func main() {
	//-----ENTRADA Y SALIDA DE DATOS-----//
	edad := 26
	name := "Sebas"
	fmt.Println("Hola mundo")            //println para nueva linea
	fmt.Printf("Mi edad es: %d\n", edad) //Formatea de acuerdo con un especificador de formato %d,%s,%t
	fmt.Printf("Mi nombre es: %s y mi edad es %d\n", name, edad)

	bandera := true
	fmt.Printf("Esto es booleano: %t\n", bandera)
	numfloat := 23.45
	fmt.Printf("El valor es: %f\n", numfloat)

	var edad1 int //Entrada enteros
	fmt.Println("Cual es tu edad:")
	fmt.Scanf("%d\n", &edad1) //Agregar & antes de la variable
	fmt.Printf("Tu edad es: %d\n", edad1)

	var nombre string //Entrada strings
	fmt.Println("Cual es tu nombre:")
	fmt.Scanf("%s\n", &nombre)
	fmt.Printf("Tu nombre es: %s\n", nombre)
}
