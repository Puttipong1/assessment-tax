package service

import (
	"fmt"
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
	level2BaseTax         = 150000.0
	level2TaxRate         = 0.1
	level3BaseTax         = 500000.0
	level3TaxRate         = 0.15
	level4BaseTax         = 1000000.0
	level4TaxRate         = 0.2
	level5BaseTax         = 2000000.0
	level5TaxRate         = 0.35
)

func (service *TaxService) CalculateTax(t request.Tax, deduction model.Deduction) response.TaxSummary {
	deduction = getTotalAllowanceByType(t.Allowances, deduction)
	netIncome := calculateNetIncome(t.TotalIncome, &deduction)
	taxLevel := calcaluteTaxLevelFromIncome(netIncome)
	summary := getTaxSummary(sumTaxLevel(taxLevel), t.Wht)
	summary.TaxLevel = taxLevel
	return summary
}

func calculateNetIncome(income float64, deduction *model.Deduction) float64 {
	return income - deduction.Personal - deduction.Donation
}

func calcaluteTaxLevelFromIncome(income float64) []response.TaxLevel {
	if income <= 150000 {
		return response.NewTaxLevel1()
	} else if income <= 500000 {
		return response.NewTaxLevel2(calculateTax(income, level2BaseTax, level2TaxRate))
	} else if income <= 1000000 {
		return response.NewTaxLevel3(calculateTax(income, level3BaseTax, level3TaxRate))
	} else if income <= 2000000 {
		return response.NewTaxLevel4(calculateTax(income, level4BaseTax, level4TaxRate))
	} else {
		return response.NewTaxLevel5(calculateTax(income, level5BaseTax, level5TaxRate))
	}
}

func calculateTax(income, base, taxRate float64) float64 {
	return decimal.NewFromFloat((income - base) * taxRate).InexactFloat64()
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

func getTaxSummary(tax, wht float64) response.TaxSummary {
	fmt.Printf("tax %f\n", tax)
	tax -= wht
	fmt.Printf("tax - wht%f\n", tax)
	if tax < 0 {
		refund := math.Abs(tax)
		return response.TaxSummary{Tax: 0, TaxRefund: &refund}
	} else {
		return response.TaxSummary{Tax: tax}
	}
}

func sumTaxLevel(taxLevel []response.TaxLevel) float64 {
	sum := 0.0
	for _, level := range taxLevel {
		fmt.Printf("level %s %f\n", level.Level, level.Tax)
		sum += level.Tax
	}
	return sum
}
