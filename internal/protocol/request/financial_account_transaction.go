package request

type RegisterTransactionRequest struct {
	TransactionGroupID int
	FinancialAccountID int
	Amount             float64
	Description        *string
}


type TransferRequest struct {
    SenderAccountID   int
    ReceiverAccountID int
    Amount            float64
    Description       string
}
