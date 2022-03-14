package delivery

type Config struct {
	Cdek       *Cdek
	Dpd        *Dpd
	Pickpoint  *Pickpoint
	Pochta     *Pochta
	Redexpress *Redexpress
	Shiptor    *Shiptor
}

func DefaultConfig() *Config {
	return &Config{}
}
