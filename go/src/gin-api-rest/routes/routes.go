package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HadnleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
