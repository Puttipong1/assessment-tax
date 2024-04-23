package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
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
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/Puttipong1/assessment-tax/server/validate"
	"github.com/Puttipong1/assessment-tax/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

const calculateTaxPath = "/tax/calculations"

func taxTestSetup(test model.Test) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validate.New()
	req := httptest.NewRequest(test.HttpMethod, test.Path, test.Body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func taxCsvTestSetup(test model.Test) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = validate.New()
	req := httptest.NewRequest(test.HttpMethod, test.Path, test.Body)
	req.Header.Add(echo.HeaderContentType, test.ContentType)
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
			TotalIncome: 2600000,
			Wht:         240000,
			Allowances: []request.Allowances{
				{AllowanceType: "donation", Amount: 100000},
			},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			var res response.TaxSummary
			err := json.Unmarshal(rec.Body.Bytes(), &res)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			fmt.Println(rec.Body.String())
			assert.True(t, reflect.DeepEqual(response.TaxSummary{Tax: 224000, TaxLevel: response.NewTaxLevel5(154000.0)}, res))
		}
	})
}

func TestCalculateTaxCSV(t *testing.T) {
	t.Run("CalculateTaxCSV Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, contentType := mockCSVMultipart("taxFile", "taxFile.csv")
		c, rec := taxCsvTestSetup(model.Test{
			HttpMethod:  http.MethodPost,
			Path:        calculateTaxPath,
			Body:        body,
			ContentType: contentType,
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTaxCSV(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}

func mockDeductionsRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"type", "amount"}).
		AddRow(common.PersonalDeductionsType, 60000)
}

func mockCSVMultipart(form, fileName string) (*bytes.Buffer, string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	filePart, _ := writer.CreateFormFile(form, fileName)
	filePart.Write([]byte("Hello, World!"))
	return body, writer.FormDataContentType()
}
