package main

import (
	"encoding/json"
	"net/http"
)

type listado struct {
	Titulo   string
	Cantidad int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		items := listado{"Casa del Arbol", 30}
		json.NewEncoder(w).Encode(items)
	})

	http.ListenAndServe(":8000", nil)
}
