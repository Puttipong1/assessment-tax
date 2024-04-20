package services

import (
	"time"

	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/models/response"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	Config *config.Config
}
type jwtCustomClaims struct {
	Role string `json:"role"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func NewJwtService(config *config.Config) *JwtService {
	return &JwtService{
		Config: config,
	}
}

func (jwtService *JwtService) CreateToken() (response.Token, int) {
	log := config.Logger()
	accessClaims := &jwtCustomClaims{
		"Admin",
		"access",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * jwtService.AccessToeknExpire())),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	// Generate encoded token and send it as response.
	accessTokenString, err := accessToken.SignedString(jwtService.JwtSignedKey())
	if err != nil {
		log.Error().Msgf("Can't signed access token : %s", err.Error())
		return response.Token{}, 0
	}
	refreshClaims := &jwtCustomClaims{
		"Admin",
		"refresh",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * jwtService.RefreshTokenExpire())),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refeshTokenString, err := refreshToken.SignedString(jwtService.JwtSignedKey())
	if err != nil {
		log.Error().Msgf("Can't signed refresh token : %s", err.Error())
		return response.Token{}, 0
	}
	token := response.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refeshTokenString,
	}
	return token, 0
}

func (jwtService *JwtService) AccessToeknExpire() time.Duration {
	return time.Duration(jwtService.Config.ServerConfig().AccessTokenExpire())
}

func (jwtService *JwtService) RefreshTokenExpire() time.Duration {
	return time.Duration(jwtService.Config.ServerConfig().RefreshTokenExpire())
}
func (JwtService *JwtService) JwtSignedKey() []byte {
	return []byte(JwtService.Config.ServerConfig().JwtSignedKey())
}
