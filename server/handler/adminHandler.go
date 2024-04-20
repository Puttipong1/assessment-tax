package handler

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/service"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService *service.AdminService
}

func NewAdminHandler(server *server.Server) *AdminHandler {
	return &AdminHandler{
		AdminService: service.NewAdminService(server),
	}
}

func (handler *AdminHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
