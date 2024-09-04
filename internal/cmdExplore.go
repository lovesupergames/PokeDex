package internal

import (
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
)

func CommandExplore(cfg *Config, cache *pokecache.Cache, location string) error {

	pokemonList, err := cfg.PokeapiClient.ListPokemon(location, cache)
	if err != nil {
		return err
	}
	for _, name := range pokemonList.PokemonEncounters {
		fmt.Println(name.Pokemon.Name)
	}
	return nil
}
