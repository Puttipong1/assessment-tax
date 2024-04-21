package common

const (
	// Deductions Type
	PersonalDeductions = "personal"
	KReceiptDeductions = "k-receipt"
	// Error Message
	BadRequestErrorMessage     = "Data sent to the server has an error or exceeds a limit"
	GetEnvErrorMessage         = "Can't get %s from env: %s"
	ShutDownServerMessage      = "Shutting down server"
	IncorrectDeductionsMessage = "%s has incorrect amount"
)
