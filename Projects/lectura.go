package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile(".//prueba.txt")

	if err != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Println(file_data)
}
