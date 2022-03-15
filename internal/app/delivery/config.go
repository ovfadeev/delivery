package delivery

import (
	"delivery/internal/app/delivery/cdek"
	"delivery/internal/app/delivery/dpd"
	"delivery/internal/app/delivery/pickpoint"
	"delivery/internal/app/delivery/pochta"
	"delivery/internal/app/delivery/redexpress"
	"delivery/internal/app/delivery/shiptor"
)

type Config struct {
	Cdek       *cdek.Cdek
	Dpd        *dpd.Dpd
	Pickpoint  *pickpoint.Pickpoint
	Pochta     *pochta.Pochta
	Redexpress *redexpress.Redexpress
	Shiptor    *shiptor.Shiptor
}

func DefaultConfig() *Config {
	return &Config{
		Cdek:       &cdek.Cdek{},
		Dpd:        &dpd.Dpd{},
		Pickpoint:  &pickpoint.Pickpoint{},
		Pochta:     &pochta.Pochta{},
		Redexpress: &redexpress.Redexpress{},
		Shiptor:    &shiptor.Shiptor{},
	}
}
