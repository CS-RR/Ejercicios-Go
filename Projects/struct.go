package main

import (
	"fmt"
)

type User struct {
	nombre, apellido string
	age              int
	bankbalance      float32
}

func main() {

	st1 := User{
		nombre:      "pepinillo",
		apellido:    "lerolero",
		age:         10,
		bankbalance: 12.34,
	}
	fmt.Println(st1)
}
