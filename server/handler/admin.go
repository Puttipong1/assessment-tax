package handler

import (
	"net/http"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/validate"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	DB *db.DB
}

func NewAdminHandler(server *server.Server) *AdminHandler {
	return &AdminHandler{DB: server.DB}
}

func (handler *AdminHandler) UpdateKReceiptDeduction(c echo.Context) error {
	log := config.Logger()
	deductions := &request.KReceiptDeductions{}
	if err := c.Bind(deductions); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(deductions); err != nil {
		return c.JSON(http.StatusBadRequest, validate.ErrorMessage(err))
	}
	err := handler.DB.UpdateDeductions(common.KReceiptDeductionsType, deductions.Amount)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, response.KReceiptDeductions{KReceipt: deductions.Amount})
}

func (handler *AdminHandler) UpdatePersonalDeduction(c echo.Context) error {
	log := config.Logger()
	deductions := &request.PersonalDeductions{}
	if err := c.Bind(deductions); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(deductions); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, validate.ErrorMessage(err))
	}
	err := handler.DB.UpdateDeductions(common.PersonalDeductionsType, deductions.Amount)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, response.PersonalDeductions{PersonalDeduction: deductions.Amount})
}
