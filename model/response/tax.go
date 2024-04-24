package response

import (
	"github.com/Puttipong1/assessment-tax/common"
	"github.com/shopspring/decimal"
)

type TaxSummary struct {
	Tax       decimal.Decimal  `json:"tax"`
	TaxRefund *decimal.Decimal `json:"taxRefund,omitempty"`
	TaxLevel  []TaxLevel       `json:"taxLevel"`
}

type TaxLevel struct {
	Level string          `json:"level"`
	Tax   decimal.Decimal `json:"tax"`
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
	zero := decimal.NewFromFloat(0.0)
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: zero},
		{Level: common.TaxLevel2, Tax: zero},
		{Level: common.TaxLevel3, Tax: zero},
		{Level: common.TaxLevel4, Tax: zero},
		{Level: common.TaxLevel5, Tax: zero},
	}
}

func NewTaxLevel2(tax decimal.Decimal) []TaxLevel {
	zero := decimal.NewFromFloat(0.0)
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: zero},
		{Level: common.TaxLevel2, Tax: tax},
		{Level: common.TaxLevel3, Tax: zero},
		{Level: common.TaxLevel4, Tax: zero},
		{Level: common.TaxLevel5, Tax: zero},
	}
}

func NewTaxLevel3(tax decimal.Decimal) []TaxLevel {
	zero := decimal.NewFromFloat(0.0)
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: zero},
		{Level: common.TaxLevel2, Tax: decimal.NewFromFloat(common.TaxLevel2Value)},
		{Level: common.TaxLevel3, Tax: tax},
		{Level: common.TaxLevel4, Tax: zero},
		{Level: common.TaxLevel5, Tax: zero},
	}
}

func NewTaxLevel4(tax decimal.Decimal) []TaxLevel {
	zero := decimal.NewFromFloat(0.0)
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: zero},
		{Level: common.TaxLevel2, Tax: decimal.NewFromFloat(common.TaxLevel2Value)},
		{Level: common.TaxLevel3, Tax: decimal.NewFromFloat(common.TaxLevel3Value)},
		{Level: common.TaxLevel4, Tax: tax},
		{Level: common.TaxLevel5, Tax: zero},
	}
}

func NewTaxLevel5(tax decimal.Decimal) []TaxLevel {
	return []TaxLevel{
		{Level: common.TaxLevel1, Tax: decimal.NewFromFloat(0.0)},
		{Level: common.TaxLevel2, Tax: decimal.NewFromFloat(common.TaxLevel2Value)},
		{Level: common.TaxLevel3, Tax: decimal.NewFromFloat(common.TaxLevel3Value)},
		{Level: common.TaxLevel4, Tax: decimal.NewFromFloat(common.TaxLevel4Value)},
		{Level: common.TaxLevel5, Tax: tax},
	}
}
