package service

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/model"
	"github.com/Puttipong1/assessment-tax/model/request"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/jszwec/csvutil"
	"github.com/shopspring/decimal"
)

type TaxService struct {
}

func NewTaxService() *TaxService {
	return &TaxService{}
}

var (
	zeroDecimal           = decimal.NewFromInt(0)
	donationMaxDeductions = decimal.NewFromFloat(100000.0)
	level2BaseTax         = decimal.NewFromFloat(150000.0)
	level2TaxRate         = decimal.NewFromFloat(0.1)
	level3BaseTax         = decimal.NewFromFloat(500000.0)
	level3TaxRate         = decimal.NewFromFloat(0.15)
	level4BaseTax         = decimal.NewFromFloat(1000000.0)
	level4TaxRate         = decimal.NewFromFloat(0.2)
	level5BaseTax         = decimal.NewFromFloat(2000000.0)
	level5TaxRate         = decimal.NewFromFloat(0.35)
)

func (service *TaxService) CalculateTax(t request.Tax, deduction model.Deduction) response.TaxSummary {
	totalIncome := decimal.NewFromFloat(t.TotalIncome)
	wht := decimal.NewFromFloat(t.Wht)
	deduction = getTotalAllowanceByType(t.Allowances, deduction)
	netIncome := calculateNetIncome(totalIncome, &deduction)
	taxLevel := calcaluteTaxLevelFromIncome(netIncome)
	summary := getTaxSummary(sumTaxLevel(taxLevel), wht)
	summary.TaxLevel = taxLevel
	return summary
}

func (service *TaxService) CalculateTaxCSV(file multipart.File, deduction model.Deduction) (*response.Tax, error) {
	taxes, err := readTaxCSV(file)
	taxCSV := []response.TaxCsv{}
	if err != nil {
		return nil, err
	}
	for _, tax := range taxes {
		res := service.CalculateTax(tax, deduction)
		taxCSV = append(taxCSV, response.TaxCsv{
			TotalIncome: decimal.NewFromFloat(tax.TotalIncome), Tax: res.Tax, TaxRefund: res.TaxRefund,
		})
	}
	return &response.Tax{Taxes: taxCSV}, nil
}

func readTaxCSV(file multipart.File) ([]request.Tax, error) {
	log := config.Logger()
	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, file); err != nil {
		log.Error().Msg(err.Error())
		return nil, &response.Error{Message: common.InvalidCsvFileMessage}
	}
	taxCsv := []model.TaxCSV{}
	if err := csvutil.Unmarshal(buffer.Bytes(), &taxCsv); err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}
	return validateTaxCSV(taxCsv)
}

func calculateNetIncome(income decimal.Decimal, deduction *model.Deduction) decimal.Decimal {
	return income.Sub(deduction.Personal).Sub(deduction.Donation)
}

func calcaluteTaxLevelFromIncome(income decimal.Decimal) []response.TaxLevel {
	if income.LessThanOrEqual(level2BaseTax) {
		return response.NewTaxLevel1()
	} else if income.LessThanOrEqual(level3BaseTax) {
		return response.NewTaxLevel2(calculateTax(income, level2BaseTax, level2TaxRate))
	} else if income.LessThanOrEqual(level4BaseTax) {
		return response.NewTaxLevel3(calculateTax(income, level3BaseTax, level3TaxRate))
	} else if income.LessThanOrEqual(level5BaseTax) {
		return response.NewTaxLevel4(calculateTax(income, level4BaseTax, level4TaxRate))
	} else {
		return response.NewTaxLevel5(calculateTax(income, level5BaseTax, level5TaxRate))
	}
}

func calculateTax(income, base, taxRate decimal.Decimal) decimal.Decimal {
	return income.Sub(base).Mul(taxRate)
}

func getTotalAllowanceByType(allowances []request.Allowances, deduction model.Deduction) model.Deduction {
	totalDonations := decimal.NewFromInt(0.0)
	for _, allowance := range allowances {
		switch allowance.AllowanceType {
		case common.DonationsDeductionsType:
			totalDonations = totalDonations.Add(decimal.NewFromFloat(allowance.Amount))
		}
	}
	deduction.Donation = decimal.Min(totalDonations, donationMaxDeductions)
	return deduction
}

func getTaxSummary(tax, wht decimal.Decimal) response.TaxSummary {
	refund := tax.Sub(wht).Abs()
	if refund.LessThan(zeroDecimal) {
		return response.TaxSummary{Tax: zeroDecimal, TaxRefund: &refund}
	} else {
		return response.TaxSummary{Tax: refund}
	}
}

func sumTaxLevel(taxLevel []response.TaxLevel) decimal.Decimal {
	sum := decimal.NewFromInt(0)
	for _, level := range taxLevel {
		sum = sum.Add(level.Tax)
	}
	return sum
}

func validateTaxCSV(csv []model.TaxCSV) ([]request.Tax, error) {
	tax := []request.Tax{}
	for i, c := range csv {
		if c.TotalIncome < c.Wht {
			return nil, &response.Error{Message: fmt.Sprintf(common.WHTIsMoreThanTotalIncomeMessage, i+2)}
		}
		tax = append(tax, request.Tax{
			TotalIncome: c.TotalIncome,
			Wht:         c.Wht,
			Allowances: []request.Allowances{
				{AllowanceType: common.DonationsDeductionsType, Amount: c.Donation},
			},
		})
	}
	return tax, nil
}
