package enum

import "fmt"

type FinancialAccountStatus uint

const (
	Verified FinancialAccountStatus = iota
	Unverified
)

// IsValid checks if the FinancialAccountStatus value is valid
func (s FinancialAccountStatus) IsValid() bool {
	switch s {
	case Verified, Unverified:
		return true
	default:
		return false
	}
}

// String converts the FinancialAccountStatus to its string representation
func (s FinancialAccountStatus) String() string {
	switch s {
	case Verified:
		return "Verified"
	case Unverified:
		return "Unverified"
	default:
		return fmt.Sprintf("Unknown(%d)", s)
	}
}
