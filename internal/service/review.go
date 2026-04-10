package service //package com as regras de negócio

import(
	"errors"
	"fmt"
	m "pizzaria/internal/models"
)

func ValidateReviewRating(review *m.Review) error{
	fmt.Println("review ", review, review.Rating >=1 && review.Rating <= 5, review.Rating <1 && review.Rating > 5)
	if review.Rating < 1 || review.Rating > 5{

		return errors.New("Favor dar uma nota de 1 a 5")
	}
	return nil
}