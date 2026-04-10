package main

import (
	"fmt"
	data "pizzaria/internal/data"
	handler "pizzaria/internal/handler"

	gin "github.com/gin-gonic/gin"
)

// var pizzas []m.Pizza

func main() {	
	// p1 := m.NewPizza(1, "Marguerita", 32.99)
	// p2 := m.NewPizza(2, "Portuguesa", 42.99)
	// p3 := m.NewPizza(3, "Camarão com Catupiry", 52.99)

	// pizzas = []m.Pizza{*p1, *p2, *p3}

	data.LoadPizzas()
	fmt.Println(data.Pizzas)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/pizzas", handler.GetPizzas)

	router.POST("/pizzas", handler.PostPizzas)

	router.GET("/pizzas/:id", handler.GetPizzasByID)

	router.DELETE("/pizzas/:id", handler.DeletePizzaByID)

	router.PUT("/pizzas/:id", handler.UpdatePizzaByID)

	router.POST("pizzas/:id/reviews", handler.PostReview)

	router.Run()
}


