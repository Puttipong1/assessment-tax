package request

type Tax struct {
	TotalIncome float64      `json:"totalIncome" validate:"gte=0.0"`
	Wht         float64      `json:"wht" validate:"gte=0.0"`
	Allowances  []Allowances `json:"allowance" validate:"required,dive,required"`
}
type Allowances struct {
	AllowanceType string  `json:"totalIncome" validate:"required"`
	Amount        float64 `json:"amount" validate:"gte=0.0"`
}
