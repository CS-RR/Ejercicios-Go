package main

import "fmt"

type st1 struct { //Asignamos atributos a st1
	nombre string
	edad   int
}

type st2 struct { //creamos una 2da estructura
	nombre string
}

type st3 struct { //Asiganmos los atributos de la estructura st1 y st2 (atributos anonimos)
	st1
	st2
}

func main() {

	nuevo := st3{st1{"onix", 15}, st2{"xion"}} //la estructura st3 contiene a la estructura 1 y 2
	fmt.Println(nuevo.st1.nombre)              //Debemos elegir que estructura tomar y que valor
	fmt.Println(nuevo.edad)
	fmt.Println(nuevo.st2.nombre) //Escogemos las 2da estructura

}
