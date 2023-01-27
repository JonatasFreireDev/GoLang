package controllers

import (
	"criando-api-rest-simplificada/models"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"Api diz": "E ai " + nome + ", tudo suave ?",
	})
}
