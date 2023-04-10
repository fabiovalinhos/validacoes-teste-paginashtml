package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiovalinhos/validacoes-teste-paginashtml/controllers"
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/database"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()

	return rotas
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

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, resposta.Code)

	fmt.Println(resposta.Body)
}
