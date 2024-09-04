package internal

import (
	"errors"
	"fmt"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
)

func CommandMapf(cfg *Config, cache *pokecache.Cache, arg string) error {
	locationsResp, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationsResp.Next
	cfg.PrevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func CommandMapb(cfg *Config, cache *pokecache.Cache, arg string) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL, cache)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
