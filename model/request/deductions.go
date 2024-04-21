package request

type KReceiptDeductions struct {
	Amount float64 `json:"amount" validate:"gt=0.0,lte=100000.0"`
}
type PersonalDeductions struct {
	Amount float64 `json:"amount" validate:"gt=10000.0,lte=100000.0"`
}
