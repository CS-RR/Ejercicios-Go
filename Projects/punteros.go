package main

import "fmt"

func main() {
	/*
			Pasos para utilizar un puntero son:
		    1.Declararlo.
		    2.Asignarle la dirección de otra variable.
		    3.Acceder a la variable desde el apuntador.
	*/
	var x, y *int
	entero := 5

	x = &entero //& Apunta a direccion de memoria
	y = &entero

	*x = 6 //*Modificamos valor de x

	fmt.Println(*x) //Se alteran las 2 variables
	fmt.Println(*y) //Se imprime el valor que esta en esa direccion de memoria
}
