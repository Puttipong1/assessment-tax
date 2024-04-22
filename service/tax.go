package service

import (
	"math"

	"github.com/Puttipong1/assessment-tax/common"
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

const (
	donationMaxDeductions = 100000.0
)

func (service *TaxService) CalculateTax(t request.Tax, deduction model.Deduction) response.TaxSummary {
	deduction = getTotalAllowanceByType(t.Allowances, deduction)
	netIncome := calculateNetIncome(t.TotalIncome, &deduction)
	tax := calcaluteTotalTaxFromIncome(netIncome)
	tax = tax - t.Wht
	return getTaxSummary(tax)
}

func calculateNetIncome(income float64, deduction *model.Deduction) float64 {
	return income - deduction.Personal - deduction.Donation
}

func calcaluteTotalTaxFromIncome(income float64) float64 {
	total := decimal.NewFromFloat(0.0)
	if income >= 150001 {
		total = total.Add(calculateTax(math.Min(income, 500000), 150000, 0.1))
	}
	if income >= 500001 {
		total = total.Add(calculateTax(math.Min(income, 1000000), 500000, 0.15))
	}
	if income >= 1000001 {
		total = total.Add(calculateTax(math.Min(income, 2000000), 10000000, 0.20))
	}
	if income >= 2000000 {
		total = total.Add(calculateTax(income, 2000000, 0.35))
	}
	return total.InexactFloat64()
}

func calculateTax(tax, base, taxRate float64) decimal.Decimal {
	return decimal.NewFromFloat((tax - base) * taxRate)
}

func getTotalAllowanceByType(allowances []request.Allowances, deduction model.Deduction) model.Deduction {
	totalDonations := 0.0
	for _, allowance := range allowances {
		switch allowance.AllowanceType {
		case common.DonationsDeductionsType:
			totalDonations += allowance.Amount
		}
	}
	deduction.Donation = math.Min(totalDonations, donationMaxDeductions)
	return deduction
}

func getTaxSummary(tax float64) response.TaxSummary {
	if tax < 0 {
		refund := math.Abs(tax)
		return response.TaxSummary{Tax: 0, TaxRefund: &refund}
	} else {
		return response.TaxSummary{Tax: tax}
	}
}
