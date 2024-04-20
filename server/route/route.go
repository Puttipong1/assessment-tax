package route

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *server.Server) {
	adminHandler := handler.NewAdminHandler(server)
	log := config.Logger()
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(middleware.RequestID())
	server.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	a := server.Echo.Group("/admin")
	a.Use(middleware.BasicAuth(func(username, password string, ctx echo.Context) (bool, error) {
		log.Info().Msgf("Test %s %s", username, password)
		if username == server.Config.AdminConfig().UserName() && password == server.Config.AdminConfig().Password() {
			return true, nil
		}
		return false, nil
	}))
	a.POST("/deductions/k-receipt", adminHandler.Login)
}
