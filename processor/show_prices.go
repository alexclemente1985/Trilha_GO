package processor

import (
	"buscador/internal/models"
	"fmt"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) { //channel operator antes do chan - LEITURA | após: escrita
	var totalPrice float64
	countPrice := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrice++
		avgPrice := totalPrice / countPrice
		fmt.Printf("[%s]Preço recebido de %s | R$ %.2f | Preço média até agora %.2f \n",price.Timestamp.Format("02-Jan-2006 15:04:05"), price.StoreName, price.Value, avgPrice)
	}

	done <- true

}
