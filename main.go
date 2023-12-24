package main

import (
	"time"

	"github.com/nollidnosnhoj/pokedexcli/internal/pokeapi"
)

type config struct {
	client pokeapi.Client
	next *string
	previous *string
}


func main() {
	cfg := config {
		client: pokeapi.NewClient(5 * time.Minute),
	}

	start(&cfg)
}