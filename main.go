package main

import (
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/database"
	"github.com/fabiovalinhos/validacoes-teste-paginashtml/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
