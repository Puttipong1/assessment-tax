package config

type Config struct {
	admin  AdminConfig
	db     DBConfig
	server ServerConfig
}

func NewConfig() *Config {
	return &Config{
		admin:  loadAdminConfig(),
		db:     LoadDBConfig(),
		server: LoadServerConfig(),
	}
}

func (config *Config) AdminConfig() *AdminConfig {
	return &config.admin
}

func (config *Config) DBConfig() *DBConfig {
	return &config.db
}

func (config *Config) ServerConfig() *ServerConfig {
	return &config.server
}
