package fetcher

import (
	"math/rand"
	"time"
)

func FetchPricesFromSite1() float64 {
	time.Sleep(1* time.Second)
	return rand.Float64()*100
}


func FetchPricesFromSite2() float64 {
	time.Sleep(2* time.Second)
	return rand.Float64()*100
}


func FetchPricesFromSite3() float64 {
	time.Sleep(3* time.Second)
	return rand.Float64()*100
}