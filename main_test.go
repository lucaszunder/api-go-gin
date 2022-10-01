package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucaszunder/api-go-gin/controller"
	"github.com/lucaszunder/api-go-gin/database"
	"github.com/lucaszunder/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func setupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func createStudentMock() {
	student := models.Student{Name: "Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func deleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCode(t *testing.T) {
	r := setupTestRoutes()
	r.GET("/message/:name", controller.WriteMessage)
	request, _ := http.NewRequest("GET", "/message/Lucas", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	responseMock := `{"message":"Olá Lucas, seja bem-vindo ao curso de Go"}`
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.Code, "Esperado era %d, recebido foi %d", http.StatusOK, response.Code)
	assert.Equal(t, responseMock, string(responseBody))
}

func TestListAllStudents(t *testing.T) {
	database.ConnectDatabase()
	createStudentMock()
	defer deleteStudentMock()
	r := setupTestRoutes()
	r.GET("/alunos", controller.ListAllStudents)
	request, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Esperado era %d, recebido foi %d", http.StatusOK, response.Code)
}

func TestGetStudentByCPF(t *testing.T) {
	database.ConnectDatabase()
	createStudentMock()
	defer deleteStudentMock()
	r := setupTestRoutes()
	r.GET("/alunos/cpf/:cpf", controller.GetStudentByCpf)
	request, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Esperado era %d, recebido foi %d", http.StatusOK, response.Code)
}

func TestGetStudentById(t *testing.T) {
	database.ConnectDatabase()
	createStudentMock()
	defer deleteStudentMock()
	r := setupTestRoutes()
	r.GET("/alunos/:id", controller.GetStudentById)

	requestPath := "/alunos/" + strconv.Itoa(ID)

	request, _ := http.NewRequest("GET", requestPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Teste", studentMock.Name)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectDatabase()
	createStudentMock()
	defer deleteStudentMock()
	r := setupTestRoutes()
	r.PATCH("/alunos/:id", controller.UpdateStudent)

	requestPath := "/alunos/" + strconv.Itoa(ID)
	requestMockBody := models.Student{Name: "Teste", CPF: "99999999999", RG: "999999999"}

	jsonValue, _ := json.Marshal(requestMockBody)
	request, _ := http.NewRequest("PATCH", requestPath, bytes.NewBuffer(jsonValue))

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "99999999999", studentMock.CPF)
	assert.Equal(t, http.StatusOK, response.Code, "Esperado era %d, recebido foi %d", http.StatusOK, response.Code)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDatabase()
	createStudentMock()
	r := setupTestRoutes()
	r.DELETE("/alunos/:id", controller.DeleteStudent)
	requestPath := "/alunos/" + strconv.Itoa(ID)

	request, _ := http.NewRequest("DELETE", requestPath, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.Code, "Esperado era %d, recebido foi %d", http.StatusOK, response.Code)
	responseMock := `{"data":"Aluno deletado com sucesso"}`
	assert.Equal(t, responseMock, string(responseBody))
}
