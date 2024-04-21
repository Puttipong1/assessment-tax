package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/Puttipong1/assessment-tax/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const calculateTaxPath = "/tax/calculations"

func taxTestSetup(test model.Test) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	req := httptest.NewRequest(test.HttpMethod, test.Path, bytes.NewBuffer(test.Json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestCalculateTax(t *testing.T) {
	t.Run("CalculateTax Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 500000.0,
			Wht:         0.0,
			Allowances: []request.Allowances{
				{AllowanceType: "donation", Amount: 0.0},
			},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Json:       body,
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			var res response.TaxSummary
			err := json.Unmarshal(rec.Body.Bytes(), &res)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.True(t, reflect.DeepEqual(response.TaxSummary{Tax: 29000}, res))
		}
	})
}

func mockDeductionsRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"type", "amount"}).
		AddRow(common.PersonalDeductionsType, 60000)
}
