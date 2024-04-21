package response

type KReceiptDeductions struct {
	KReceipt float64 `json:"kReceipt" validate:"gt=0.0,lte=100000.0"`
}
type PersonalDeductions struct {
	PersonalDeduction float64 `json:"personalDeduction" validate:"gt=10000.0,lte=100000.0"`
}
