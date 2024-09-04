package internal

import (
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
)

func CommandPokedex(cfg *Config, cache *pokecache.Cache, location string) error {
	for name, _ := range cfg.CaughtPokemon {
		fmt.Printf("-%s\n", name)
	}
	return nil
}
