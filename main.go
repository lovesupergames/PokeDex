package main

import (
	"github.com/lovesupergames/PokeDex/internal"
	"github.com/lovesupergames/PokeDex/internal/pokeAPI"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"time"
)

func main() {
	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)
	pokeClient := pokeAPI.NewClient(5 * time.Second)
	cfg := &internal.Config{
		CaughtPokemon: map[string]pokeAPI.Pokemon{},
		PokeapiClient: pokeClient,
	}

	internal.StartRepl(cfg, cache)
}
