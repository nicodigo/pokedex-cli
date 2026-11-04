package pokeapi

import (
	"net/http"
	"time"

	"github.com/nicodigo/pokedex-cli/internal/pokecache"
)

type PokeApiClient struct{
	httpClient *http.Client
	cache *pokecache.Cache
}

func NewPokeApiClient(cacheInterval time.Duration) *PokeApiClient {
	return &PokeApiClient{
		httpClient: &http.Client{},
		cache: pokecache.NewCache(cacheInterval),
	}
}
