package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiovalinhos/validacoes-teste-paginashtml/controllers"
	"github.com/gin-gonic/gin"
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

	if resposta.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado era %d", resposta.Code, http.StatusOK)
	}
}
