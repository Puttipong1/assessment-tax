package response

type TaxSummary struct {
	Tax       float64  `json:"tax"`
	taxRefund *float64 `json:"taxRefund,omitempty"`
}
