package main

import (
	"alunos/database"
	"alunos/routes"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context){
	c.JSON(200, gin.H{
		"id": "1",
		"nome": "Alex",
	})
}

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequest()
}