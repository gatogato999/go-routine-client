package main

import (
	"context"
	"fmt"
	"time"
)

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Temp string `json:"temp"`
}

func main() {
	cntx := context.Background()
	cuncurr(cntx, time.Millisecond*8)
}

func cuncurr(contxt context.Context, timeout time.Duration) {
	cancelContext, cancel := context.WithTimeout(contxt, timeout)
	defer cancel()

	start := time.Now()

	cities2check := []string{"city1", "city2"}

	resultChanal := make(chan string)
	for _, cityName := range cities2check {
		go FetchCityDataconcurrently(cityName, resultChanal)
	}

	for range cities2check {
		select {
		case res := <-resultChanal:
			fmt.Println(res)
		case <-cancelContext.Done():
			fmt.Printf("operation canceled : it took more than %v\n", timeout)
			return
		}
	}
	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}
