package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(httpTimeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: httpTimeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
