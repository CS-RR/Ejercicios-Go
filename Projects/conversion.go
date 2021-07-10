package main

import (
	"fmt"
	"strconv" //Para convertir int a str
)
func main() {

	//-----CONVERSION DE TIPOS-----//
	edad := 26
	edad_str := strconv.Itoa(edad) //Int a str
	edad2 := "23"
	edad_int, _ := strconv.Atoi(edad2) //str a int. Se usa _ pues la funcion Atoi retorna 2 valores.

	fmt.Println("Mi edad es: " + edad_str)
	fmt.Println(edad_int + 10) //Se realiza la suma de enteros

}
