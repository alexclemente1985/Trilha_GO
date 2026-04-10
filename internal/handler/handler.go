package handler

import (
	"net/http"
    "strconv"
	m "pizzaria/internal/models"
	s "pizzaria/internal/service"
	"pizzaria/internal/data"
	"github.com/gin-gonic/gin"
)

func PostReview(c *gin.Context) {
	pizzaID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
        return
    }

	var newReview m.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := s.ValidateReviewRating(&newReview); err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	for i, p := range data.Pizzas {
        if p.ID == pizzaID {
            p.Review = append(p.Review, newReview)
            data.Pizzas[i] = p
            data.SavePizza()
            c.JSON(http.StatusCreated, p)
            return
        }
    }

	c.JSON(http.StatusNotFound, gin.H{"message": "Pizza not found"})

}