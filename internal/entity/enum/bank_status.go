package enum

type BankStatus uint

const (
	BankInactive BankStatus = iota
	BankActive
)