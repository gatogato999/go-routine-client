package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Temp string `json:"temp"`
}

func main() {
	seq()
}

func seq() {
	start := time.Now()

	cities2check := []string{"city1", "city2"}

	for _, cityName := range cities2check {
		fetchCityDataSequentialy(cityName)
	}

	fmt.Printf("all done in %d milliseconds\n", time.Since(start).Milliseconds())
}

func fetchCityDataSequentialy(cityName string) {
	sub_start := time.Now()

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
	fmt.Printf("%s took %d milliseconds\n", cityName, time.Since(sub_start).Milliseconds())
	fmt.Print("-----------------\n\n")
}
