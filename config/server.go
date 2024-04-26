package config

import (
	"os"
)

type ServerConfig struct {
	port string
}

func LoadServerConfig() ServerConfig {
	return ServerConfig{
		port: ":" + os.Getenv("PORT"),
	}
}

func (serverConfig *ServerConfig) Port() string {
	return serverConfig.port
}
