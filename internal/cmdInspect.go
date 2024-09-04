package internal

import (
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
)

func CommandInspect(cfg *Config, cache *pokecache.Cache, pokemonName string) error {

	PokemonStruct, ok := cfg.CaughtPokemon[pokemonName]
	if ok {
		fmt.Printf("Name:%s\n", PokemonStruct.Name)
		fmt.Printf("Height%d\n", PokemonStruct.Height)
		fmt.Printf("Weight%d\n", PokemonStruct.Weight)
		fmt.Println("Stats:")
		for _, stat := range PokemonStruct.Stats {
			fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeObj := range PokemonStruct.Types {
			fmt.Printf("- %s\n", typeObj.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon yet\n")
	}
	return nil
}
