package response

type TaxSummary struct {
	Tax       float64  `json:"tax"`
	TaxRefund *float64 `json:"taxRefund,omitempty"`
}
