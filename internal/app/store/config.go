package store

type Config struct {
	DBUrl string
}

func DefaultConfig() *Config {
	return &Config{}
}
