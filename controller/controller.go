package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucaszunder/api-go-gin/database"
	"github.com/lucaszunder/api-go-gin/models"
)

func ListAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"Error": err.Error(),
		})
		return
	}
	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func UpdateStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func GetStudentByCpf(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}
