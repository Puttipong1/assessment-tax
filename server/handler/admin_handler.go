package handler

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/models/request"
	"github.com/Puttipong1/assessment-tax/models/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/services"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminService *services.AdminService
}

func NewAdminHandler(server *server.Server) *AdminHandler {
	return &AdminHandler{
		AdminService: services.NewAdminService(server),
	}
}

func (handler *AdminHandler) Login(c echo.Context) error {
	log := config.Logger()
	login := request.Login{}
	if err := c.Bind(&login); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(login); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}

	token, status := handler.AdminService.Login(login)
	if status != 0 {
		return c.NoContent(status)
	}
	return c.JSON(http.StatusOK, token)
}
