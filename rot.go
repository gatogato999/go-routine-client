package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func FetchCityDataconcurrently(cityName string, ch chan<- string) {
	subStart := time.Now()

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
			time.Since(subStart).Milliseconds())
	} else {
		ch <- fmt.Sprintf("%s doesn't exist\t took %d milliseconds\n\n", cityName, time.Since(subStart).Milliseconds())
	}
}

func fetchCityDataSequentialy(cityName string) {
	subStart := time.Now()

	url := fmt.Sprintf("http://localhost:3000/cities?name=%s", cityName)
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	resBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		log.Println(err)
		return
	}

	var data []City
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("-----------------")
	fmt.Printf("%s city --> %s degree\t\t", data[0].Name, data[0].Temp)
	fmt.Printf("%s took %d milliseconds\n", cityName, time.Since(subStart).Milliseconds())
	fmt.Print("-----------------\n\n")
}
