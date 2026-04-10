package models


type Pizza struct {
	ID    int     `json:"id"` // Público no Go, minúsculo no JSON
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
	Review []Review `json:"reviews"`
}

func NewPizza(
	ID int,
	Nome string,
	Preco float64) *Pizza {
	return &Pizza{ID: ID, Nome: Nome, Preco: Preco}
}
