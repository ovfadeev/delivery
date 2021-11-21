package store

type Config struct {
	DbUrl string
}

func DefaultConfig() *Config {
	return &Config{}
}
