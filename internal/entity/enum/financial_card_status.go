package enum

type FinancialCardStatus uint

const (
	Active FinancialCardStatus = iota
	Inactive
	Lost
	Stolen
)

// IsValid checks if the given FinancialCardStatus is valid.
func IsValidFinancialCardStatus(s FinancialCardStatus) bool {
	switch s {
	case Active, Inactive, Lost, Stolen:
		return true
	default:
		return false
	}
}
