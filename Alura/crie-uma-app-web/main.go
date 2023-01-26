package main

import (
	"crie-um-app-web/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProtudos()

	// produtos := []Produto{
	// 	{"Camiseta", "Azul", 49, 10},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// }

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
