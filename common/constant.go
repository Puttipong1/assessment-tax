package common

const (
	// Deductions
	PersonalDeductions      = "Personal Deductions"
	KReceiptDeductions      = "K-Receipt Deductions"
	PersonalDeductionsType  = "personal"
	KReceiptDeductionsType  = "k-receipt"
	DonationsDeductionsType = "donation"
	// Error Message
	BadRequestErrorMessage     = "Data sent to the server has an error or exceeds a limit"
	GetEnvErrorMessage         = "Can't get %s from env: %s"
	ShutDownServerMessage      = "Shutting down server"
	IncorrectDeductionsMessage = "%s has incorrect amount"
)
