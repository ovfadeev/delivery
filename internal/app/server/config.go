package server

import (
	"delivery/internal/app/delivery"
	"delivery/internal/app/store"
)

type Config struct {
	ServerAddr string
	LogLevel   string
	DBUrl      string
	Store      *store.Config
	Delivery   *delivery.Config
}

func DefaultConfig() *Config {
	return &Config{
		ServerAddr: ":8080",
		LogLevel:   "info",
		DBUrl:      "",
		Store:      store.DefaultConfig(),
		Delivery:   delivery.DefaultConfig(),
	}
}
