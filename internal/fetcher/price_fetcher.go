package fetcher

import (
	"buscador/internal/models"
	"math/rand"
	"sync"
	"time"
)

func FetchPrices (priceChannel chan<- models.PriceDetail){// func FetchPrices(priceChannel chan<- float64) { //para acrescentar dados no channel informado
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done() // garante que não ocorra o deadlock (é executado quando tudo o que estiver neste escopo for executado. Vai informar que o processo acabou)
		//price1 = fetcher.FetchPricesFromSite1()
		priceChannel <- FetchPricesFromSite1() // <- channel operator
	}()

	go func() {
		defer wg.Done()
		//price2 = FetchPricesFromSite2()
		priceChannel <- FetchPricesFromSite2()
	}()

	go func() {
		defer wg.Done()
		//price3 = FetchPricesFromSite3()
		priceChannel <- FetchPricesFromSite3()
	}()

	go func () {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPricesFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail {
		StoreName: "A",
		Value: rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPricesFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail {
		StoreName: "B",
		Value: rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPricesFromSite3() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail {
		StoreName: "C",
		Value: rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(3 * time.Second)
	prices := []float64{
		rand.Float64()*100,
		rand.Float64()*100,
		rand.Float64()*100,
//		rand.Float64()*100,
	}

	for _, price := range prices {
		priceChannel <- models.PriceDetail {
		StoreName: "D",
		Value: price,
		Timestamp: time.Now(),
	}
	}
}