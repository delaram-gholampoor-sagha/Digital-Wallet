package enum

type FinancialCardType uint

const (
	Debit FinancialCardType = iota
	Credit
	Gift
)

// IsValid checks if the given FinancialCardType is valid.
func IsValidFinancialCardType(t FinancialCardType) bool {
	switch t {
	case Debit, Credit, Gift:
		return true
	default:
		return false
	}
}
