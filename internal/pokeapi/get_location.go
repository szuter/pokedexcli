package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(pageUrl *string) (Lokation, error) {
	url := baseURL + "location-area/" + *pageUrl
	if cached, ok := c.cache.Get(url); ok {
		location := Lokation{}
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return Lokation{}, err
		}
		return location, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Lokation{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Lokation{}, err
	}
	if res.StatusCode > 299 {
		return Lokation{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Lokation{}, err
	}
	location := Lokation{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Lokation{}, err
	}
	c.cache.Add(url, data)
	return location, nil
}
