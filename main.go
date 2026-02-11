package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	cities2check := []string{"city1", "city2", "city4"}

	for _, cityName := range cities2check {
		fetchCityData(cityName)
	}

	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}

func fetchCityData(cityName string) {
}
