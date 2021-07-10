package main

import (
	"fmt"
)
func main() {
	//Ciclo For
	for i := 0; i < 10; i++ {
		fmt.Println("Hola mundo")
		fmt.Println(i)
	}
	//Ciclo While
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	//Uso del break
	x := 0
	for {
		fmt.Println(x)
		x++
		if x > 10 {
			break
		}
	}
	//Ciclo continuo
	y := 0
	for {
		if y == 2 {
			fmt.Println("Se omitió 2 y continuó")
			y++
			continue
		}
		fmt.Println(y)
		y++
		if y > 10 {
			break
		}
	}
}
