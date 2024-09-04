package internal

import (
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"math/rand"
)

func CommandCatch(cfg *Config, cache *pokecache.Cache, pokemonName string) error {
	pokemon, err := cfg.PokeapiClient.ListPokemonName(pokemonName, cache)
	if err != nil {
		return err
	}
	name := pokemon.Name

	fmt.Printf("Throwing a Pokeball at %s\n", name)
	//catcher

	baseExp := pokemon.BaseExperience
	catchChance := 80
	if baseExp > 100 {
		catchChance = 50
	}
	if catchChance < rand.Intn(101) {
		fmt.Printf("Catched Pokemon: %s\n", name)
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped\n", name)
	}
	return nil
}
