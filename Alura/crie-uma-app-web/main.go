package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul", 49, 10},
		{"Teclado", "RGB", 799, 49},
		{"Teclado", "RGB", 799, 49},
		{"Teclado", "RGB", 799, 49},
		{"Teclado", "RGB", 799, 49},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
