package enum

type FinancialAccountTransactionStatus uint

const (
	Peniding FinancialAccountTransactionStatus = iota
	Completed
	Failed
	Reversed
	OnHold
	Cancelled
)
