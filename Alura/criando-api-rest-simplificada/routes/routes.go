package routes

import (
	"criando-api-rest-simplificada/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.GetAllStudents)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.PostStudent)
	r.GET("/alunos/:id", controllers.GetByIdStudent)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.PutStudent)
	r.GET("/alunos/cpf/:cpf", controllers.GetByCPF)

	r.Run()
}
