package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Temp string `json:"temp"`
}

func main() {
	start := time.Now()

	cities2check := os.Args[1:]

	resultChanal := make(chan string)
	for _, cityName := range cities2check {
		go fetchCityData(cityName, resultChanal)
	}
	for range cities2check {
		fmt.Println(<-resultChanal)
	}

	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}

func fetchCityData(cityName string, ch chan<- string) {
	sub_start := time.Now()

	url := fmt.Sprintf("http://localhost:3000/cities?name=%s", cityName)
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	var data []City
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprintf("%s city --> %s degree\t\t  took %d milliseconds\n",
		data[0].Name,
		data[0].Temp,
		time.Since(sub_start).Milliseconds())
}
