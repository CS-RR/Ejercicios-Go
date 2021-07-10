package main

import "fmt"

type user struct {
	nombre, apellido string
	edad             int
}

func (usuario user) nombre_completo() string { //Usa un identificador (usuario,this)
	return usuario.nombre + " " + usuario.apellido //Une los datos de la estructura
}

/*func (this *user) set_name(n string)  { //Copiamos una referencia no la estructura completa *
	this.nombre=n
}*/

func main() {

	st1 := new(user)
	st1.nombre = "Onix" //asignamos valores a los datos de la estructura
	st1.apellido = "Aguila"
	//st1.set_name("Xino")//Asignamos nuevo valor de la estructura
	fmt.Println(st1.nombre_completo()) //imprimimos la funcion con los identificadores
}
