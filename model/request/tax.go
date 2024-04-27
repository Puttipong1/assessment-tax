package request

type Tax struct {
	TotalIncome float64      `json:"totalIncome" validate:"gte=0.0"`
	Wht         float64      `json:"wht" validate:"gte=0.0,ltcsfield=TotalIncome"`
	Allowances  []Allowances `json:"allowances" validate:"unique=AllowanceType,dive"`
}
type Allowances struct {
	AllowanceType string  `json:"allowanceType" validate:"required,eq=donation|eq=k-receipt"`
	Amount        float64 `json:"amount" validate:"gte=0.0"`
}
