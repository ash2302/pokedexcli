package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	locationAreaName := args[0]
	locationAreaResp, err := cfg.pokeapiClient.GetLocation(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationAreaName)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
