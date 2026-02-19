package main

import (
	"encoding/json"
	"fmt"
	"io"
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
