package main

import (
	"delivery/internal/app/server"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "config/.env", "config file")

	if err := godotenv.Load(configPath); err != nil {
		log.Fatal("No config .env file found")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	config := server.DefaultConfig()
	config.ServerAddr = os.Getenv("SERVER_ADDR")
	config.LogLevel = os.Getenv("LOG_LEVEL")
	config.DBUrl = fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SSLMODE"),
	)

	s := server.NewConfig(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
