package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	mi_nombre("onix")
}

func mi_nombre(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
