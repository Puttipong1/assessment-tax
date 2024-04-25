package model

type TaxCSV struct {
	TotalIncome float64 `csv:"totalIncome"`
	Wht         float64 `csv:"wht"`
	Donation    float64 `csv:"donation"`
}
