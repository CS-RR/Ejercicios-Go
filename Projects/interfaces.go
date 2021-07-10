package main

import (
	"fmt"
)

type usuario interface { //Se crea la intefaz que se usará en las estructuras
	permisos() int
	nombre() string
}

/*
Si los metodos de la interfaz  estan implementados en la estructura
la estructura esta implementando los atributos de la interfaz
*/
type admin struct { //Creamos nuestra estructura le asignamos un atributo para identificar
	name string
}

func (this admin) permisos() int { //nivel de permiso que tiene nuestra estructura
	return 5
}
func (this admin) nombre() string { //nombre que se le asigna a la estructura
	return this.name
}

/*
Agregamos otro usuario/estructura  para mostrar que no tiene permisos de administrador
*/
type auxiliar struct { //cambiamos el nombre de la estrucura
	name string
}

func (this auxiliar) permisos() int { //nivel de permiso que tiene la estrucura
	return 3
}
func (this auxiliar) nombre() string { //nombre que se le asigna a la estructura
	return this.name
}

//Funcion para autenticar usuario o administrador
func autenticador(user usuario) string { //creamos una funcion donde se revisan los permisos de cada usuario
	if user.permisos() > 4 { //si el usuario/estructura tiene permiso mayor a 4 es un admin
		return user.nombre() + " Tiene permisos de administrador"
	}
	return user.nombre() + " No tiene permisos de administrador" //si no es administrador
}
func main() {

	aux1 := admin{"Onix"}    //Le asignamos un nombre a la 1era estructura para identificarlo
	aux2 := auxiliar{"Xion"} //Le asignamos un nombre a la 2da estructura para identificarlo
	fmt.Println(autenticador(aux1))
	fmt.Println(autenticador(aux2))

	//Se puede aplicar un arreglo para tomar las estructuras
	usuarios := []usuario{admin{"Onix"}, auxiliar{"Xion"}}
	for _, usuario := range usuarios { //Se crea un ciclo para recorrer cada estructura.
		fmt.Println(autenticador(usuario))
	}
}
