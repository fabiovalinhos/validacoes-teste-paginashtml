package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiovalinhos/validacoes-teste-paginashtml/controllers"
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/database"
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	// Simplifica as mensagens de teste
	gin.SetMode(gin.ReleaseMode)

	///
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "43215678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeDaSaudacao(t *testing.T) {

	r := SetupDasRotasDeTeste()

	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/bruce", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	// irei usar o testify
	// if resposta.Code != http.StatusOK {
	// 	t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	// }

	assert := assert.New(t)
	assert.Equal(http.StatusOK, resposta.Code, "O valor recebido poderia ser igual ao esperado")

	mockDaResposta := `{"API diz:":"E ai bruce, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	assert.Equal(mockDaResposta, string(respostaBody))

	fmt.Println(string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, resposta.Code)

	fmt.Println(resposta.Body)
}

func TestBuscaAlunoPorCPF(t *testing.T) {

	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/43215678901", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
