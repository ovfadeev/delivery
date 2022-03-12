package main

import (
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
	// CDEK
	if os.Getenv("CDEK") == "on" {
		config.Delivery.Cdek.URL = os.Getenv("CDEK_URL")
		config.Delivery.Cdek.LOGIN = os.Getenv("CDEK_LOGIN")
	}
	// DPD
	if os.Getenv("DPD") == "on" {
		config.Delivery.Dpd.URL = os.Getenv("DPD_URL")
		config.Delivery.Dpd.CLIENT = os.Getenv("DPD_CLIENT")
		config.Delivery.Dpd.KEY = os.Getenv("DPD_KEY")
	}
	// PICKPOINT
	if os.Getenv("PICKPOINT") == "on" {
		config.Delivery.Pickpoint.URL = os.Getenv("PICKPOINT_URL")
		config.Delivery.Pickpoint.LOGIN = os.Getenv("PICKPOINT_LOGIN")
		config.Delivery.Pickpoint.PASS = os.Getenv("PICKPOINT_PASS")
		config.Delivery.Pickpoint.IKN = os.Getenv("PICKPOINT_IKN")
	}
	// POCHTA RF
	if os.Getenv("POCHTA") == "on" {
		config.Delivery.Pochta.URL = os.Getenv("POCHTA_URL")
		config.Delivery.Pochta.ACCOUNT = os.Getenv("POCHTA_ACCOUNT")
	}
	// REDEXPRESS
	if os.Getenv("REDEXPRESS") == "on" {
		config.Delivery.Redexpress.URL = os.Getenv("REDEXPRESS_URL")
		config.Delivery.Redexpress.LOGIN = os.Getenv("REDEXPRESS_LOGIN")
		config.Delivery.Redexpress.PASS = os.Getenv("REDEXPRESS_PASS")
	}
	// SHIPTOR
	if os.Getenv("SHIPTOR") == "on" {
		config.Delivery.Shiptor.URL = os.Getenv("SHIPTOR_URL")
		config.Delivery.Shiptor.KEY = os.Getenv("SHIPTOR_KEY")
	}

	// server start
	s := server.NewConfig(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
