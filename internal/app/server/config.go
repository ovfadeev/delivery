package server

import "delivery/internal/app/store"

type Config struct {
	ServerAddr string
	LogLevel   string
	Store      *store.Config
}

func DefaultConfig() *Config {
	return &Config{
		ServerAddr: ":8080",
		LogLevel:   "info",
		Store:      store.DefaultConfig(),
	}
}
