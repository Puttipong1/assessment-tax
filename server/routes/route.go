package routes

import (
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureRoutes(server *server.Server) {
	adminHandler := handler.NewAdminHandler(server)
	server.Echo.Use(middleware.Recover())
	server.Echo.POST("/api/login", adminHandler.Login)
}
