package config

type Config struct {
	Auth   AdminConfig
	DB     DBConfig
	Server ServerConfig
}

func NewConfig() *Config {
	return &Config{
		Auth:   loadAdminConfig(),
		DB:     LoadDBConfig(),
		Server: LoadServerConfig(),
	}
}
