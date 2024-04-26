package handler_test

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const calculateTaxPath = "/tax/calculations"

var (
	CorrectCSV = `totalIncome,wht,donation
500000.0,0.0,0.0
600000.0,40000.0,20000.0
750000.0,50000.0,15000.0`
)

func taxTestSetup(test model.Test) (echo.Context, *httptest.ResponseRecorder) {
	decimal.MarshalJSONWithoutQuotes = true
	e := echo.New()
	e.Validator = validate.New()
	req := httptest.NewRequest(test.HttpMethod, test.Path, test.Body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func taxCsvTestSetup(test model.Test) (echo.Context, *httptest.ResponseRecorder) {
	decimal.MarshalJSONWithoutQuotes = true
	e := echo.New()
	e.Validator = validate.New()
	req := httptest.NewRequest(test.HttpMethod, test.Path, test.Body)
	req.Header.Add(echo.HeaderContentType, test.ContentType)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestCalculateTax(t *testing.T) {
	t.Run("CalculateTax Level 1 Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 150000,
			Wht:         10000,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: 20000},
				{AllowanceType: common.KReceiptDeductionsType, Amount: 20000}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(0), TaxLevel: response.NewTaxLevel1()})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax Level 2 Success", func(t *testing.T) {
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
				{AllowanceType: common.DonationsDeductionsType, Amount: 200000.0}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(19000), TaxLevel: response.NewTaxLevel2(decimal.NewFromInt(19000))})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax Level 3 Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 800000.0,
			Wht:         40000.0,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: 100000},
				{AllowanceType: common.KReceiptDeductionsType, Amount: 100000}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(8500), TaxLevel: response.NewTaxLevel3(decimal.NewFromInt(13500))})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax Level 4 Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 1500000.0,
			Wht:         100000.0,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: 100000},
				{AllowanceType: common.KReceiptDeductionsType, Amount: 100000}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(68000), TaxLevel: response.NewTaxLevel4(decimal.NewFromInt(58000))})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax Level 5 Success", func(t *testing.T) {
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
				{AllowanceType: common.DonationsDeductionsType, Amount: 100000},
				{AllowanceType: common.KReceiptDeductionsType, Amount: 100000}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(206500), TaxLevel: response.NewTaxLevel5(decimal.NewFromInt(136500))})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax with tax refund Success", func(t *testing.T) {
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		mock.ExpectQuery("SELECT").WillReturnRows(mockDeductionsRows())
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 2600000,
			Wht:         500000,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: 100000},
				{AllowanceType: common.KReceiptDeductionsType, Amount: 100000}},
		})
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			refund := decimal.NewFromInt(53500)
			expect, err := json.Marshal(response.TaxSummary{Tax: decimal.NewFromInt(0), TaxRefund: &refund, TaxLevel: response.NewTaxLevel5(decimal.NewFromInt(136500))})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
	t.Run("CalculateTax has incorrect total income and wht fail", func(t *testing.T) {
		database, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		body, _ := json.Marshal(request.Tax{
			TotalIncome: -44444,
			Wht:         -44444,
		},
		)
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("CalculateTax wht is more than total income fail", func(t *testing.T) {
		database, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 500000,
			Wht:         600000,
		},
		)
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("CalculateTax has incorrect allowances.amount fail", func(t *testing.T) {
		database, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 500000,
			Wht:         200000,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: -5555},
			},
		},
		)
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("CalculateTax has incorrect allowances.allowanceType fail", func(t *testing.T) {
		database, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 500000,
			Wht:         200000,
			Allowances: []request.Allowances{
				{AllowanceType: common.GetEnvErrorMessage, Amount: 30000},
			},
		},
		)
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("CalculateTax has same allowances.allowanceType fail", func(t *testing.T) {
		database, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer database.Close()
		body, _ := json.Marshal(request.Tax{
			TotalIncome: 500000,
			Wht:         200000,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: 30000},
				{AllowanceType: common.DonationsDeductionsType, Amount: 30000},
			},
		},
		)
		c, rec := taxTestSetup(model.Test{
			HttpMethod: http.MethodPost,
			Path:       calculateTaxPath,
			Body:       bytes.NewBuffer(body),
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTax(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
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
		body, contentType := mockCSVMultipart("taxFile", "taxFile.csv", CorrectCSV)
		c, rec := taxCsvTestSetup(model.Test{
			HttpMethod:  http.MethodPost,
			Path:        calculateTaxPath,
			Body:        body,
			ContentType: contentType,
		})
		h := &handler.TaxHandler{DB: &db.DB{DB: database}, TaxService: service.NewTaxService()}
		if assert.NoError(t, h.CalculateTaxCSV(c)) {
			expect, err := json.Marshal(response.Tax{Taxes: []response.TaxCsv{
				{TotalIncome: decimal.NewFromInt(500000), Tax: decimal.NewFromInt(29000)},
				{TotalIncome: decimal.NewFromInt(600000), Tax: decimal.NewFromInt(2000)},
				{TotalIncome: decimal.NewFromInt(750000), Tax: decimal.NewFromInt(11250)},
			}})
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, strings.Replace(rec.Body.String(), "\n", "", 1), string(expect))
		}
	})
}

func mockDeductionsRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{"type", "amount"}).
		AddRow(common.PersonalDeductionsType, 60000).
		AddRow(common.KReceiptDeductionsType, 50000).
		AddRow(common.DonationsDeductionsType, 100000)
}

func mockCSVMultipart(form, fileName, csv string) (*bytes.Buffer, string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()
	filePart, _ := writer.CreateFormFile(form, fileName)
	filePart.Write([]byte(csv))
	return body, writer.FormDataContentType()
}
