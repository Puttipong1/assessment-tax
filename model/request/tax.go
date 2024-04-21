package request

type Tax struct {
	TotalIncome float64      `json:"totalIncome" validate:"required,gte=0.0"`
	Wht         float64      `json:"wht" validate:"required,gte=0.0"`
	Allowances  []Allowances `json:"allowance" validate:"required,dived,required"`
}
type Allowances struct {
	AllowanceType string  `json:"totalIncome" validate:"required,gte=0.0"`
	Amount        float64 `json:"amount" validate:"required,gte=0.0"`
}
