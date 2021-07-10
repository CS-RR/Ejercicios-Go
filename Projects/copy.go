package main

import "fmt"

func main() {

	//copy: (destino,fuente)
	slice := []int{1, 2, 3, 4}
	//cad2 := make([]int,5)//tiene capacidad para copiar 5 elementos si no usa todos rellena con 0
	cad2 := make([]int, len(slice))//usar len para asegurar copiar todos los elementos

	copy(cad2, slice)

	fmt.Println(slice)
	fmt.Print(cad2)
}
