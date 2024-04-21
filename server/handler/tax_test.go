package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const calculateTaxPath = "/tax/calculations"

func taxTestSetup(test model.Test, t *testing.T) (echo.Context, *httptest.ResponseRecorder, sqlmock.Sqlmock, *handler.TaxHandler) {
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer database.Close()
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	req := httptest.NewRequest(test.HttpMethod, test.Path, bytes.NewBuffer(test.Json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec, mock, &handler.TaxHandler{DB: &db.DB{DB: database}}
}
func TestCalculateTax(t *testing.T) {
	t.Run("CalculateTax Success", func(t *testing.T) {
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 0.0,
			Wht:         0.0,
			Allowances: []request.Allowances{
				{AllowanceType: "donation", Amount: 0.0},
			},
		})
		c, rec, _, handler := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Json:       body,
		}, t)
		if assert.NoError(t, handler.CalculateTax(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
