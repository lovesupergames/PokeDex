package internal

import (
	"github.com/lovesupergames/PokeDex/internal/pokecache"
	"os"
)

func CommandExit(cfg *Config, cache *pokecache.Cache, arg string) error {
	os.Exit(0)
	return nil
}
