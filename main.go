package main

import (
	"fmt"
	"time"
)

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Temp string `json:"temp"`
}

func main() {
}

func seq() {
	start := time.Now()

	cities2check := []string{"city1", "city2"}

	for _, cityName := range cities2check {
		fetchCityDataSequentialy(cityName)
	}

	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}

func cuncurr() {
	start := time.Now()

	cities2check := []string{"city1", "city2"}

	resultChanal := make(chan string)
	for _, cityName := range cities2check {
		go FetchCityDataconcurrently(cityName, resultChanal)
	}

	for range cities2check {
		fmt.Println(<-resultChanal)
	}
	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}
