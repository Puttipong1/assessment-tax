package common

const (
	// CSV
	TaxCsvFileName  = "taxFile.csv"
	TaxCsvFieldName = "taxFile"
	// Tax Level
	TaxLevel1      = "0-150000"
	TaxLevel2      = "150,001-500,000"
	TaxLevel2Value = 35000.0
	TaxLevel3      = "500,001-1,000,000"
	TaxLevel3Value = 75000.0
	TaxLevel4      = "1,000,001-2,000,000"
	TaxLevel4Value = 200000.0
	TaxLevel5      = "2,000,001 ขึ้นไป"
	// Deductions
	PersonalDeductions      = "Personal Deductions"
	KReceiptDeductions      = "K-Receipt Deductions"
	PersonalDeductionsType  = "personal"
	KReceiptDeductionsType  = "k-receipt"
	DonationsDeductionsType = "donation"
	// Error Message
	BadRequestErrorMessage             = "Data sent to the server has an error or exceeds a limit"
	GetEnvErrorMessage                 = "Can't get %s from env: %s"
	ShutDownServerMessage              = "Shutdown server"
	IncorrectDeductionsMessage         = "%s has incorrect amount"
	InvalidCsvFileMessage              = "CSV File is incorrect or corrupt"
	CSVWHTIsMoreThanTotalIncomeMessage = "Line %d: Total income should be more than Wht"
	CSVHasLowerThanZeroMessage         = "Line %d: Has lower than 0.0"
)
