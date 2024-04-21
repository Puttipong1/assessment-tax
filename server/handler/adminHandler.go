package handler

import (
	"fmt"
	"net/http"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	db *db.DB
}

func NewAdminHandler(server *server.Server) *AdminHandler {
	return &AdminHandler{db: server.DB}
}

func (handler *AdminHandler) UpdateKReceipt(c echo.Context) error {
	log := config.Logger()
	deductions := &request.KReceiptDeductions{}
	if err := c.Bind(deductions); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(deductions); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: fmt.Sprintf(common.IncorrectDeductionsMessage, "K-Receipt Deductions")})
	}
	err := handler.db.UpdateDeductions(common.KReceiptDeductions, deductions.Amount)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, response.KReceiptDeductions{KReceipt: deductions.Amount})
}
