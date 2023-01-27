package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 2", Historia: "Historia 2"},
	}

	routes.HandleRequest()
}
