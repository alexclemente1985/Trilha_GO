package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var price1, price2, price3 float64
	var wg sync.WaitGroup
	wg.Add(3)

	go func () {
		defer wg.Done() // garante que não ocorra o deadlock (é executado quando tudo o que estiver neste escopo for executado. Vai informar que o processo acabou)
		price1 = fetcher.FetchPricesFromSite1()
	}()

	go func () {
		defer wg.Done()
		price2 = fetcher.FetchPricesFromSite2()
	}()

	go func () {
		defer wg.Done()
		price3 = fetcher.FetchPricesFromSite3()
	}()

	wg.Wait()


    fmt.Printf("R$ %.2f \n", price1)
    fmt.Printf("R$ %.2f \n", price2)
    fmt.Printf("R$ %.2f \n", price3)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}