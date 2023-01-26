package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
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

func conectaComBD() *sql.DB {
	conexao := "user=postgres dbname=loja password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBD()
	defer db.Close()

	selectAllProducts, err := db.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	// produtos := []Produto{
	// 	{"Camiseta", "Azul", 49, 10},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// 	{"Teclado", "RGB", 799, 49},
	// }

	temp.ExecuteTemplate(w, "Index", produtos)
}
