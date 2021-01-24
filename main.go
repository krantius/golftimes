package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	api := &API{
		cli: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	times, err := api.GetTimes(MilesSquare, "01-20-2021")
	if err != nil {
		panic(err)
	}

	fmt.Println(times)
}
