package enum

type CardTransactionStatus uint

const (
	Pending CardTransactionStatus = iota
	Completedd
	Failedd
	Reversedd
	OnHoldd
	Cancelledd
)
