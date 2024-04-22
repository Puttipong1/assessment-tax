package service

import (
	"math"

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

func (service *TaxService) CalculateTax(t request.Tax, deduction model.Deduction) response.TaxSummary {
	netIncome := t.TotalIncome - deduction.Personal
	baseTax, taxRate := getTaxRate(netIncome)
	total := decimal.NewFromFloat((netIncome - baseTax) * taxRate)
	tax := total.InexactFloat64() - t.Wht
	return getTaxSummary(tax)
}

func getTaxRate(total float64) (float64, float64) {
	if total < 150000.99 {
		return 0, 0
	} else if total >= 150001 && total <= 500000.99 {
		return 150000, 0.1
	} else if total >= 500001 && total <= 1000000.99 {
		return 500000, 0.15
	} else if total >= 1000001 && total <= 2000000.99 {
		return 1000000, 0.2
	} else {
		return 2000000, 0.35
	}
}

func getTaxSummary(tax float64) response.TaxSummary {
	if tax < 0 {
		refund := math.Abs(tax)
		return response.TaxSummary{Tax: 0, TaxRefund: &refund}
	} else {
		return response.TaxSummary{Tax: tax}
	}

}
