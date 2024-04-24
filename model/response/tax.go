package response

import (
	"github.com/Puttipong1/assessment-tax/common"
	"github.com/shopspring/decimal"
)

type TaxSummary struct {
	Tax       float64    `json:"tax"`
	TaxRefund *float64   `json:"taxRefund,omitempty"`
	TaxLevel  []TaxLevel `json:"taxLevel"`
}

type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type Tax struct {
	Taxes []TaxCsv `json:taxes`
}
type TaxCsv struct {
	TotalIncome decimal.Decimal `json:"totalIncome"`
	Tax         decimal.Decimal `json:"tax"`
	TaxRefund   decimal.Decimal `json:"taxRefund,omitempty"`
}

func NewTaxLevel1() []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: 0.0},
		{Level: common.TaxLevel2, Tax: 0.0},
		{Level: common.TaxLevel3, Tax: 0.0},
		{Level: common.TaxLevel4, Tax: 0.0},
		{Level: common.TaxLevel5, Tax: 0.0},
	}
}

func NewTaxLevel2(tax float64) []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: 0.0},
		{Level: common.TaxLevel2, Tax: tax},
		{Level: common.TaxLevel3, Tax: 0.0},
		{Level: common.TaxLevel4, Tax: 0.0},
		{Level: common.TaxLevel5, Tax: 0.0},
	}
}

func NewTaxLevel3(tax float64) []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: 0.0},
		{Level: common.TaxLevel2, Tax: common.TaxLevel2Value},
		{Level: common.TaxLevel3, Tax: tax},
		{Level: common.TaxLevel4, Tax: 0.0},
		{Level: common.TaxLevel5, Tax: 0.0},
	}
}

func NewTaxLevel4(tax float64) []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: 0.0},
		{Level: common.TaxLevel2, Tax: common.TaxLevel2Value},
		{Level: common.TaxLevel3, Tax: common.TaxLevel3Value},
		{Level: common.TaxLevel4, Tax: tax},
		{Level: common.TaxLevel5, Tax: 0.0},
	}
}

func NewTaxLevel5(tax float64) []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: 0.0},
		{Level: common.TaxLevel2, Tax: common.TaxLevel2Value},
		{Level: common.TaxLevel3, Tax: common.TaxLevel3Value},
		{Level: common.TaxLevel4, Tax: common.TaxLevel4Value},
		{Level: common.TaxLevel5, Tax: tax},
	}
}
