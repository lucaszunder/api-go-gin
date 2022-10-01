package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaszunder/api-go-gin/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/message/:name", controller.WriteMessage)
	r.GET("/alunos", controller.ListAllStudents)
	r.POST("/alunos", controller.CreateStudent)
	r.GET("/alunos/:id", controller.GetStudentById)
	r.GET("/alunos/cpf/:cpf", controller.GetStudentByCpf)
	r.DELETE("/alunos/:id", controller.DeleteStudent)
	r.PATCH("/alunos/:id", controller.UpdateStudent)
	r.Run()
}
