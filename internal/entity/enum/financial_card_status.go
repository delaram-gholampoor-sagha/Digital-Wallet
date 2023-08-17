package enum

type FinancialCardStatus uint

const (
	Active FinancialCardStatus = iota
	Inactive
	Lost
	Stolen
)
