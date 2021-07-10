package main

import (
	"fmt"
)

func main() {
	//-----CONDICIONALES-----//
/*
	== Igua a
	!= Diferente de
	< menor que
	> mayor que
	>= mayor igual
	<= menor igual
	&& AND toma en cuenta las 2 variables
	|| OR  Compara 2 condiciones diferentes
*/
  x := 24
  y := 45
  if x > y {
	  fmt.Printf("%d es mayor que %d\n", x, y)
  } else if x == y {
	  fmt.Printf("%d es igual a %d\n", x, y)
  } else {
	  fmt.Printf("%d es menor que %d\n", x, y)
  }
}