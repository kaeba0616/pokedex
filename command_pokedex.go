package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("you must provide no arg")
	}
	fmt.Println("Your Pokedex:")
	pokedex := cfg.pokedex
	if len(pokedex) < 1 {
		fmt.Println("you have no pokemon. dude")
	}
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
