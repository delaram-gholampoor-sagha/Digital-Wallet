package enum

type AccountTransactionStatus uint

const (
	Peniding AccountTransactionStatus = iota
	Completed
	Failed
	Reversed
	OnHold
	Cancelled
)
