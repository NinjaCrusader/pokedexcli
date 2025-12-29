package main

import (
	"math/rand"
	"time"

	"github.com/NinjaCrusader/pokedexcli/internal"
	"github.com/NinjaCrusader/pokedexcli/internal/pokecache"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cfg := &config{
		Next:     nil,
		Previous: nil,
		Cache:    pokecache.NewCache(10 * time.Minute),
		Pokedex:  make(map[string]internal.Pokemon),
	}

	startRepl(cfg)
}
