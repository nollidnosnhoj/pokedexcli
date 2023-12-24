package pokeapi

import (
	"net/http"
	"time"

	"github.com/nollidnosnhoj/pokedexcli/internal/pokecache"
)

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}