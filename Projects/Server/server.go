package main

import (
	"net/http"
)

func main() {
	//Lectura de archivos estaticos
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.ListenAndServe(":8000", nil)
}
