package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile(".//Prueba.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Println(file_data)
}
