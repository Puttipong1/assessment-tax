package server

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	Server struct {
		Echo   *echo.Echo
		Config *config.Config
		DB     *db.DB
	}

	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func NewServer() *Server {
	echo := echo.New()
	echo.Validator = &CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	config := config.NewConfig()
	return &Server{
		Echo:   echo,
		Config: config,
		DB:     db.Init(config.DBConfig()),
	}
}
