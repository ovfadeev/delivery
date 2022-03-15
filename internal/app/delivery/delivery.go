package delivery

import (
	"net/url"
)

type Delivery struct {
	config *Config
}

func (d *Delivery) NewConfig(config *Config) *Delivery {
	return &Delivery{
		config: config,
	}
}

func (d *Delivery) GetPoints(query url.Values) ([]byte, error) {

	return []byte(""), nil
}

func (d *Delivery) GetCourier(query url.Values) ([]byte, error) {

	return []byte(""), nil
}
