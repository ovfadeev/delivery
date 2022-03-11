package api

import "delivery/internal/app/delivery"

type Config struct {
	Cdek       *delivery.Cdek
	Dpd        *delivery.Dpd
	Pickpoint  *delivery.Pickpoint
	Pochta     *delivery.Pochta
	Redexpress *delivery.Redexpress
	Shiptor    *delivery.Shiptor
}

func DefaultConfig() *Config {
	return &Config{}
}
