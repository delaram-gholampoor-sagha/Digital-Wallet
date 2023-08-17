package enum

type FinancialCardTransactionStatus uint

const (
	Pending FinancialCardTransactionStatus = iota
	Completedd
	Failedd
	Reversedd
	OnHoldd
	Cancelledd
)
