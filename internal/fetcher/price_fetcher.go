package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- float64) { //para acrescentar dados no channel informado
	var wg sync.WaitGroup
	wg.Add(3)
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

	wg.Wait()
	close(priceChannel)
}

func FetchPricesFromSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPricesFromSite2() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

func FetchPricesFromSite3() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}
