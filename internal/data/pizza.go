package data

import (
	"fmt"
	"os"
	"encoding/json"
	m "pizzaria/internal/models"
)

var Pizzas []m.Pizza

func SavePizza(){
	file, err := os.Create("data/pizza.json")

	if err != nil {
		fmt.Println("Error file: ",err)
		return
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(Pizzas); err != nil{
		fmt.Println("Error encoding JSON: ", err)
	}
}

func LoadPizzas(){
	file, err := os.Open("data/pizza.json")

	if err != nil {
		fmt.Println("Error file: ",err)
		return
	}

	defer file.Close() //garante que o arquivo seja fechado ao se encerrar a função

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Pizzas); err != nil {
		fmt.Println("Error decoding JSON ", err)
	}
}