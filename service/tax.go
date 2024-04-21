package service

import (
	"github.com/Puttipong1/assessment-tax/model"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/shopspring/decimal"
)

type TaxService struct {
}

func NewTaxService() *TaxService {
	return &TaxService{}
}

func (service *TaxService) CalculateTax(tax request.Tax, deduction model.Deduction) response.TaxSummary {
	if tax.TotalIncome < 150000.99 {
		return response.TaxSummary{Tax: 0.0}
	} else if tax.TotalIncome >= 150001 && tax.TotalIncome <= 500000.99 {
		amount := decimal.NewFromFloat((tax.TotalIncome - 150000 - tax.Wht - deduction.Personal) * 0.1)
		return response.TaxSummary{Tax: amount.InexactFloat64()}
	} else if tax.TotalIncome >= 500001 && tax.TotalIncome <= 1000000.99 {
		amount := decimal.NewFromFloat((tax.TotalIncome - 500000 - tax.Wht - deduction.Personal) * 0.15)
		return response.TaxSummary{Tax: amount.InexactFloat64()}
	} else if tax.TotalIncome >= 1000001 && tax.TotalIncome <= 2000000.99 {
		amount := decimal.NewFromFloat((tax.TotalIncome - 150000 - tax.Wht - deduction.Personal) * 0.2)
		return response.TaxSummary{Tax: amount.InexactFloat64()}
	} else {
		amount := decimal.NewFromFloat((tax.TotalIncome - 2000000 - tax.Wht - deduction.Personal) * 0.35)
		return response.TaxSummary{Tax: amount.InexactFloat64()}
	}
}
