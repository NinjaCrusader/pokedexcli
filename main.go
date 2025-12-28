package main

import (
	"time"

	"github.com/NinjaCrusader/pokedexcli/internal/pokecache"
)

func main() {
	cfg := &config{
		Next:     nil,
		Previous: nil,
		Cache:    pokecache.NewCache(10 * time.Minute),
	}

	startRepl(cfg)
}
