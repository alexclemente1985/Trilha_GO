package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/models"
	"buscador/processor"
	"fmt"

	//"sync"
	"time"
)

func main() {
	start := time.Now()
	//var price1, price2, price3 float64
	//priceChannel := make(chan models.PriceDetail)//make(chan float64) //channel para resumir todos as goroutines em uma variável só

	priceChannel := make(chan models.PriceDetail, 7) //buffer -- len (valores no buffer) cap (tamanho máximo)

	// var wg, showWg sync.WaitGroup      //showWg -> garante a execução do último cálculo (as vezes fecha wg e não mostra)
	//var showWg sync.WaitGroup
	// wg.Add(3)
	//showWg.Add(1)
	done := make(chan bool)

	//go func() { //precisa estar antes das funções com wg.Done()
	//somente vai ser invocada no momento que o priceChannel receber cada função
	// var totalPrice float64
	// countPrice := 0.0
	// for price := range priceChannel {
	// 	totalPrice += price
	// 	countPrice++
	// 	avgPrice := totalPrice / countPrice
	// 	fmt.Printf("Preço recebido: R$ %.2f | Preço médio até agora: R$ %.2f \n", price, avgPrice)
	// }
	//defer showWg.Done()
	//processor.ShowPriceAVG(priceChannel)
	//}()

	// go func() {
	// 	defer wg.Done() // garante que não ocorra o deadlock (é executado quando tudo o que estiver neste escopo for executado. Vai informar que o processo acabou)
	// 	//price1 = fetcher.FetchPricesFromSite1()
	// 	priceChannel <- fetcher.FetchPricesFromSite1() // <- channel operator
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	//price2 = fetcher.FetchPricesFromSite2()
	// 	priceChannel <- fetcher.FetchPricesFromSite2()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	//price3 = fetcher.FetchPricesFromSite3()
	// 	priceChannel <- fetcher.FetchPricesFromSite3()
	// }()

	// wg.Wait()
	// close(priceChannel)
	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAVG(priceChannel, done)

	<-done //trava a execução ao garantir que somente prossiga o processo quando o valor de done chegar
	//showWg.Wait() //colocado aqui para garantir que só imprima algo quando o canal for concluído e fechado

	// fmt.Printf("R$ %.2f \n", price1)
	// fmt.Printf("R$ %.2f \n", price2)
	// fmt.Printf("R$ %.2f \n", price3)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
