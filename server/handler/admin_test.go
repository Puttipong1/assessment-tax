package handler_test

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Puttipong1/assessment-tax/db"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/Puttipong1/assessment-tax/server"
	"github.com/Puttipong1/assessment-tax/server/handler"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type test struct {
	httpMethod string
	path       string
	json       []byte
}

func setup(test test) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &server.CustomValidator{Validator: validator.New(validator.WithRequiredStructEnabled())}
	req := httptest.NewRequest(test.httpMethod, test.path, bytes.NewBuffer(test.json))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}
func TestUpdateKReceipt(t *testing.T) {
	t.Run("Update K-Receipt Successful", func(t *testing.T) {
		amount := 60000.0
		body, _ := json.Marshal(request.KReceiptDeductions{Amount: amount})
		c, rec := setup(test{
			httpMethod: http.MethodPost,
			path:       "/admin/deductions/k-receipt",
			json:       body,
		})
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec("UPDATE deductions").WillReturnResult(driver.RowsAffected(1))
		defer database.Close()
		h := &handler.AdminHandler{DB: &db.DB{DB: database}}
		if assert.NoError(t, h.UpdateKReceiptDeduction(c)) {
			response := response.KReceiptDeductions{}
			err := json.Unmarshal(rec.Body.Bytes(), &response)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NoError(t, err)
			assert.Equal(t, amount, response.KReceipt)
		}
	})

	t.Run("Update K-Receipt Fail (amount <= 0.0)", func(t *testing.T) {
		amount := 0.0
		body, _ := json.Marshal(request.KReceiptDeductions{Amount: amount})
		c, rec := setup(test{
			httpMethod: http.MethodPost,
			path:       "/admin/deductions/k-receipt",
			json:       body,
		})
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec("UPDATE deductions").WillReturnResult(driver.RowsAffected(1))
		defer database.Close()
		h := &handler.AdminHandler{DB: &db.DB{DB: database}}
		if assert.NoError(t, h.UpdateKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Update K-Receipt Fail (amount > 100000.0)", func(t *testing.T) {
		amount := 100000.1
		body, _ := json.Marshal(request.KReceiptDeductions{Amount: amount})
		c, rec := setup(test{
			httpMethod: http.MethodPost,
			path:       "/admin/deductions/k-receipt",
			json:       body,
		})
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec("UPDATE deductions").WillReturnResult(driver.RowsAffected(1))
		defer database.Close()
		h := &handler.AdminHandler{DB: &db.DB{DB: database}}
		if assert.NoError(t, h.UpdateKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Update K-Receipt Fail (error during update)", func(t *testing.T) {
		amount := 100000.0
		body, _ := json.Marshal(request.KReceiptDeductions{Amount: amount})
		c, rec := setup(test{
			httpMethod: http.MethodPost,
			path:       "/admin/deductions/k-receipt",
			json:       body,
		})
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec("UPDATE deductions").WillReturnError(*new(error))
		defer database.Close()
		h := &handler.AdminHandler{DB: &db.DB{DB: database}}
		if assert.NoError(t, h.UpdateKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Update K-Receipt Fail (affected row = 0)", func(t *testing.T) {
		amount := 100000.0
		body, _ := json.Marshal(request.KReceiptDeductions{Amount: amount})
		c, rec := setup(test{
			httpMethod: http.MethodPost,
			path:       "/admin/deductions/k-receipt",
			json:       body,
		})
		database, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock.ExpectExec("UPDATE deductions").WillReturnResult(driver.RowsAffected(0))
		defer database.Close()
		h := &handler.AdminHandler{DB: &db.DB{DB: database}}
		if assert.NoError(t, h.UpdateKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}
