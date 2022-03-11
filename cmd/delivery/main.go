package main

import (
	"delivery/internal/app/api"
	"delivery/internal/app/store"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

	sConfig := store.DefaultConfig()
	sConfig.DBUrl = fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SSLMODE"),
	)

	aConfig := api.DefaultConfig()

	// aConfig.Cdek.URL = os.Getenv("CDEK_URL")
	// aConfig.Cdek.KEY = os.Getenv("CDEK_KEY")

	aConfig.Shiptor.URL = os.Getenv("SHIPTOR_URL")
	aConfig.Shiptor.KEY = os.Getenv("SHIPTOR_KEY")

}
