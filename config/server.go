package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	port               string
	accessTokenExpire  int
	refreshTokenExpire int
	jwtSignedKey       string
}

func LoadServerConfig() ServerConfig {
	l := Logger()
	accessTokenExpire, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRE_IN_SECOND"))
	if err != nil {
		l.Error().Msgf("Can't get ACCESS_TOKEN_EXPIRE from env: %s", err)
		accessTokenExpire = 5
	}
	refreshTokenExpire, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRE_IN_SECOND"))
	if err != nil {
		l.Error().Msgf("Can't get REFRESH_TOKEN_EXPIRE from env: %s", err)
		refreshTokenExpire = 60
	}
	jwtSignedKey := os.Getenv("JWT_SIGNED_KEY")
	if jwtSignedKey == "" {
		l.Error().Msgf("Can't get JWT_SIGNED_KEY from env: %s", err)
		os.Exit(0)
	}
	return ServerConfig{
		port:               ":" + os.Getenv("SERVER_PORT"),
		accessTokenExpire:  accessTokenExpire,
		refreshTokenExpire: refreshTokenExpire,
		jwtSignedKey:       jwtSignedKey,
	}
}

func (serverConfig *ServerConfig) Port() string {
	return serverConfig.port
}
func (serverConfig *ServerConfig) AccessTokenExpire() int {
	return serverConfig.accessTokenExpire
}
func (serverConfig *ServerConfig) RefreshTokenExpire() int {
	return serverConfig.refreshTokenExpire
}
func (serverConfig *ServerConfig) JwtSignedKey() string {
	return serverConfig.jwtSignedKey
}
