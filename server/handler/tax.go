package handler

import (
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/labstack/echo/v4"
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
	return nil
}
