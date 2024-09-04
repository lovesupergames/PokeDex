package pokeAPI

import (
	"encoding/json"
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"io"
	"net/http"
)

// ListPokemonName -
func (c *Client) ListPokemonName(pokemonName string, cache *pokecache.Cache) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	pokecache.NewCache(cache.Interval)
	if data, ok := cache.Get(url); ok {
		pokemonList := Pokemon{}
		err := json.Unmarshal(data, &pokemonList)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonList, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	cache.Add(url, dat)

	pokemonList := Pokemon{}
	err = json.Unmarshal(dat, &pokemonList)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonList, nil
}
