package pokeAPI

import (
	"encoding/json"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(locationName string, cache *pokecache.Cache) (PokemonListJSON, error) {
	url := baseURL + "/location-area/" + locationName
	pokecache.NewCache(cache.Interval)
	if data, ok := cache.Get(url); ok {
		pokemonList := PokemonListJSON{}
		err := json.Unmarshal(data, &pokemonList)
		if err != nil {
			return PokemonListJSON{}, err
		}
		return pokemonList, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonListJSON{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonListJSON{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonListJSON{}, err
	}
	cache.Add(url, dat)

	pokemonList := PokemonListJSON{}
	err = json.Unmarshal(dat, &pokemonList)
	if err != nil {
		return PokemonListJSON{}, err
	}

	return pokemonList, nil
}
