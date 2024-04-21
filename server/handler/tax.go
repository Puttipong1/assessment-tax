package handler

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type TaxHandler struct {
	DB *db.DB
}

func NewTaxHandler(server *server.Server) *TaxHandler {
	return &TaxHandler{
		DB: server.DB,
	}
}

func (handler *TaxHandler) CalculateTax(c echo.Context) error {
	var tax = request.Tax{}
	if err := c.Bind(&tax); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(tax); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
	}
	return nil
}
