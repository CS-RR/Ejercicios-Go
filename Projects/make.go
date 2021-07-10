package main

import "fmt"

func main() {
	slice := make([]int, 3, 5) //longitud de 3
	slice = append(slice, 2)   //agrega un elemento
	fmt.Println(len(slice))    //longitud
	fmt.Println(cap(slice))    //capacidad
	fmt.Println(slice)
	fmt.Println(slice)
}
