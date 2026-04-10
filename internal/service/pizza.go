package service //package com as regras de negócio

import(
	"errors"
	m "pizzaria/internal/models"
)

func ValidatePizzaPrice(pizza *m.Pizza) error{
	if pizza.Preco < 0{
		return errors.New("O preço da pizza não pode ser negativo")
	}
	return nil
}