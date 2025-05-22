package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(cfg *config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatal("response failed StatusCode: %d \n Body: %s\n ", res.StatusCode, res.Body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	return nil
}
