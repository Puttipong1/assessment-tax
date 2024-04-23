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
	"github.com/Puttipong1/assessment-tax/service"
	"github.com/labstack/echo/v4"
)

type TaxHandler struct {
	DB         *db.DB
	TaxService *service.TaxService
}

func NewTaxHandler(server *server.Server) *TaxHandler {
	return &TaxHandler{
		DB:         server.DB,
		TaxService: service.NewTaxService(),
	}
}

func (handler *TaxHandler) CalculateTax(c echo.Context) error {
	log := config.Logger()
	var tax = request.Tax{}
	if err := c.Bind(&tax); err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: common.BadRequestErrorMessage})
	}
	if err := c.Validate(tax); err != nil {
		return c.JSON(http.StatusBadRequest, validate.ErrorMessage(err))
	}
	deduction, err := handler.DB.GetDeductions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: ""})
	}
	log.Info().Msgf("%f", handler.TaxService.CalculateTax(tax, deduction).Tax)
	return c.JSON(http.StatusOK, handler.TaxService.CalculateTax(tax, deduction))
}

func (handler *TaxHandler) CalculateTaxCSV(c echo.Context) error {
	log := config.Logger()
	csv, err := c.FormFile("taxFile")
	if err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
	}
	if csv.Filename != common.TaxCsvFileName {
		log.Error().Msgf("File name is %s expect %s", csv.Filename, common.TaxCsvFileName)
		return c.JSON(http.StatusBadRequest, response.Error{Message: "Tax Csv file name is incorrect"})
	}
	file, err := csv.Open()
	if err != nil {
		log.Error().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Message: err.Error()})
	}
	defer file.Close()
	return c.JSON(http.StatusOK, "")
}
