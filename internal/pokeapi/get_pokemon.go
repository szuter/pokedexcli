package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName *string) (Pokemon, error) {
	url := baseURL + "pokemon/" + *pokemonName
	if cached, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(cached, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	return pokemon, nil
}
