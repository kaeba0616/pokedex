package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	pokemon_name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemon_name)
	if err != nil {
		return err
	}

	base_experience := pokemon.BaseExperience
	chance := rand.Intn(base_experience)
	if chance > 40 {
		fmt.Printf("%s escape!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}
