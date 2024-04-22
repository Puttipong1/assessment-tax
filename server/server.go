package server

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/server/validate"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
	DB     *db.DB
}

func NewServer() *Server {
	echo := echo.New()
	echo.Validator = validate.New()
	config := config.NewConfig()
	return &Server{
		Echo:   echo,
		Config: config,
		DB:     db.Init(config.DBConfig()),
	}
}
