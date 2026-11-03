package pokeapi

import "net/http"

type PokeApiClient struct{
	httpClient *http.Client
}

func NewPokeApiClient() *PokeApiClient {
	return &PokeApiClient{httpClient: &http.Client{}}
}
