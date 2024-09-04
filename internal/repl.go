package internal

import (
	"bufio"
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokeAPI"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"os"
	"strings"
)

type Config struct {
	PokeapiClient    pokeAPI.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]pokeAPI.Pokemon
}

func StartRepl(cfg *Config, cache *pokecache.Cache) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandArgs := ""
		if len(words) == 2 {
			commandArgs = words[1]
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.Callback(cfg, cache, commandArgs)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, *pokecache.Cache, string) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"explore": {
			Name:        "explore",
			Description: "Get Pokemon list on location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch Pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect Pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Get all caught pokemons list",
			Callback:    CommandPokedex,
		},
	}
}
