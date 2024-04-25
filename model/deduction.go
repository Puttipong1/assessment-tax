package model

import "github.com/shopspring/decimal"

type Deduction struct {
	Personal decimal.Decimal
	Donation decimal.Decimal
	KReceipt decimal.Decimal
}
