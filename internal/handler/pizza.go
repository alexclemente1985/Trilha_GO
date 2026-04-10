package handler

import (
	"fmt"
	"strconv"
	"net/http"
	m "pizzaria/internal/models"
	"pizzaria/internal/data"
	s "pizzaria/internal/service"

	gin "github.com/gin-gonic/gin"
)

func DeletePizzaByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	for i, p := range data.Pizzas{
		if p.ID == id {
			//... operador variádico - permite descompactar a lista [i+1:] em elementos individuais, o que é necessãrio para o append funcionar (o append não trabalha com adição de um bloco de lista inteiro) 
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...) //[:i] permite percorrer a lista até o elemento i (sem envolvê-lo) | [i+1:] percorre lista da posição i+1 até o final (pula o elemento i)
			data.SavePizza()
			c.JSON(http.StatusOK, gin.H{"method": "delete"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found."})

	
}

func UpdatePizzaByID(c *gin.Context){
	idParam := c.Param("id")
	id, errID := strconv.Atoi(idParam) //converte para inteiro o parâmetro id da url
	if errID != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errID.Error()})
		return
	}

	var updatedPizza m.Pizza
	
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := s.ValidatePizzaPrice(&updatedPizza); err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	for i, p := range data.Pizzas{
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id //garante que o ID seja o mesmo
			data.SavePizza()
			c.JSON(http.StatusOK, gin.H{"message": "Pizza successfully updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found."})
}



func GetPizzasByID(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	for _, p := range data.Pizzas{
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found."})
	
}

func GetPizzas(c *gin.Context) {
	fmt.Println(data.Pizzas)
	c.JSON(http.StatusOK, gin.H{"pizzas": data.Pizzas})
}

func PostPizzas(c *gin.Context) {
	var newPizza m.Pizza

	if err := c.ShouldBindJSON(&newPizza); err != nil { //Já captura os dados do JSON enviado e verifica se está correto (SERIALIZE)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := s.ValidatePizzaPrice(&newPizza); err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	newPizza.ID = len(data.Pizzas) + 1

	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Pizza " + newPizza.Nome + " adicionada com sucesso na base de dados"})

}