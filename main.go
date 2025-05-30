package main

import (
	"time"

	"github.com/kaeba0616/pokedex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
