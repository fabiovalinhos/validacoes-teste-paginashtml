package routes

import (
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)

	r.Run()
}
