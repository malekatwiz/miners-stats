package minerstats

type Config struct {
	Operations []OpConfig `yaml:"operations"`
}

type OpConfig struct {
	Currency string `yaml:"currency"`
	PoolUrl  string `yaml:"poolUrl"`
	Wallet   string `yaml:"wallet"`
}
