package main

import (
	"fmt"
)

func main() {

	var arreglo [10]int
	arreglo2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arreglo)
	fmt.Println(arreglo2)
	dias := [7]string{"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo"}
	fmt.Println(dias)
	fmt.Println(dias[3]) //Para imprimir un element especifico
	//Arreglo bidimensional
	var tabla [2][2]int
	var tabla2 = [][]int{{1, 2}, {3, 4}}
	fmt.Println(tabla)
	fmt.Println(tabla2)
	var matriz [3][2]int
	/*
	       [ 00  01 ] Para acceder a una posicion de un arreglo bidimensional se identifica
	       [ 10  11 ] Arreglo y posicion
	   	[ 20  21 ]
	*/
	matriz[0][0] = 5
	matriz[0][1] = 6
	matriz[1][0] = 7
	matriz[1][1] = 8
	matriz[2][0] = 9
	matriz[2][1] = 10
	fmt.Println(matriz)

}
