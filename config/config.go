package config

type Config struct {
	Client ClientConfig
}

type ClientConfig struct {
	Version string
	Timeout int64
}

var globalConfig = &Config{
	Client: ClientConfig{
		Version: "v1",
		Timeout: 10,
	},
}

func GetConfig() *Config {
	return globalConfig
}
