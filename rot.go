package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Cuncurr() {
	start := time.Now()

	cities2check := []string{"city1", "city2"}

	resultChanal := make(chan string)
	for _, cityName := range cities2check {
		go FetchCityData(cityName, resultChanal)
	}

	for range cities2check {
		fmt.Println(<-resultChanal)
	}
	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}

func FetchCityData(cityName string, ch chan<- string) {
	sub_start := time.Now()

	url := fmt.Sprintf("http://localhost:3000/cities?name=%s", cityName)
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	var data []City
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	if len(data) > 0 {
		ch <- fmt.Sprintf("%s city --> %s C°\t took %d milliseconds\n\n",
			data[0].Name,
			data[0].Temp,
			time.Since(sub_start).Milliseconds())
	} else {
		ch <- fmt.Sprintf("%s doesn't exist\t took %d milliseconds\n\n", cityName, time.Since(sub_start).Milliseconds())
	}
}

func cuncurr() {
	start := time.Now()

	cities2check := []string{"city1", "city2"}

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
		ch <- fmt.Sprintln(err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	var data []City
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		ch <- fmt.Sprintln(err)
		return
	}

	if len(data) > 0 {
		ch <- fmt.Sprintf("%s city --> %s C°\t took %d milliseconds\n\n",
			data[0].Name,
			data[0].Temp,
			time.Since(sub_start).Milliseconds())
	} else {
		ch <- fmt.Sprintf("%s doesn't exist\t took %d milliseconds\n\n", cityName, time.Since(sub_start).Milliseconds())
	}
}
