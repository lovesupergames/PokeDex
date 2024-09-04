package internal

import (
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
)

func CommandHelp(cfg *Config, cache *pokecache.Cache, arg string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}
