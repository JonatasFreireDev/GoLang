package main

import (
	"criando-api-rest-simplificada/models"
	"criando-api-rest-simplificada/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Jhow", CPF: "2131321321313", RG: "545465456454"},
		{Nome: "Ana", CPF: "2324234234", RG: "1231231232"},
		{Nome: "And", CPF: "2131321321313", RG: "545465456454"},
	}
	routes.HandleRequests()
}
