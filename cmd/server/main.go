package main

import (
	"delivery/internal/app/api"
	"delivery/internal/app/server"
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

	config := server.DefaultConfig()
	config.ServerAddr = os.Getenv("SERVER_ADDR")
	config.LogLevel = os.Getenv("LOG_LEVEL")
	// DATABASE
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

	// API
	aConfig := api.DefaultConfig()

	// CDEK
	if os.Getenv("CDEK") == "on" {
		aConfig.Cdek.URL = os.Getenv("CDEK_URL")
		aConfig.Cdek.LOGIN = os.Getenv("CDEK_LOGIN")
	}
	// DPD
	if os.Getenv("DPD") == "on" {
		aConfig.Dpd.URL = os.Getenv("DPD_URL")
		aConfig.Dpd.CLIENT = os.Getenv("DPD_CLIENT")
		aConfig.Dpd.KEY = os.Getenv("DPD_KEY")
	}
	// PICKPOINT
	if os.Getenv("PICKPOINT") == "on" {
		aConfig.Pickpoint.URL = os.Getenv("PICKPOINT_URL")
		aConfig.Pickpoint.LOGIN = os.Getenv("PICKPOINT_LOGIN")
		aConfig.Pickpoint.PASS = os.Getenv("PICKPOINT_PASS")
		aConfig.Pickpoint.IKN = os.Getenv("PICKPOINT_IKN")
	}
	// POCHTA RF
	if os.Getenv("POCHTA") == "on" {
		aConfig.Pochta.URL = os.Getenv("POCHTA_URL")
		aConfig.Pochta.ACCOUNT = os.Getenv("POCHTA_ACCOUNT")
	}
	// REDEXPRESS
	if os.Getenv("REDEXPRESS") == "on" {
		aConfig.Redexpress.URL = os.Getenv("REDEXPRESS_URL")
		aConfig.Redexpress.LOGIN = os.Getenv("REDEXPRESS_LOGIN")
		aConfig.Redexpress.PASS = os.Getenv("REDEXPRESS_PASS")
	}
	// SHIPTOR
	if os.Getenv("SHIPTOR") == "on" {
		aConfig.Shiptor.URL = os.Getenv("SHIPTOR_URL")
		aConfig.Shiptor.KEY = os.Getenv("SHIPTOR_KEY")
	}
}
