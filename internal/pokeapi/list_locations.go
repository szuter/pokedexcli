package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespLocations, error) {
	url := baseURL + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if cached, ok := c.cache.Get(url); ok {
		fmt.Println("Using cache")
		location := RespLocations{}
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return RespLocations{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err

	}

	res, err := c.httpClient.Do(req)
	if res.StatusCode > 299 {
		return RespLocations{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}

	location := RespLocations{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return RespLocations{}, err
	}
	c.cache.Add(url, data)
	return location, nil
}
