package main

import (
	"fmt"
	//"strconv" //Para convertir int a str
)

func main() {

	var matriz []int
	matriz2 := []int{1, 4, 3, 6, 5, 8, 10} //Declaracion slice no tiene tamaño especifico
	fmt.Println(matriz)
	fmt.Println(matriz2)
	fmt.Println(len(matriz2)) //Longitud
	slice := matriz2[1:4]     //Partir la matriz: Se declara donde inicia y cuantas posiciones toma
	fmt.Println(slice)
	slic := matriz2[:4] //Inicia en posición 1 y Tomará hasta la posicion 4
	fmt.Println(slic)
	sli := matriz2[3:] //Inicia desde la posicion 3 hasta el final del slice
	fmt.Println(sli)

}
