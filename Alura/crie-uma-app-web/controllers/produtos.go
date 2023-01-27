package controllers

import (
	"crie-um-app-web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
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

func PostProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversao do preco: ", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversao do preco: ", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
