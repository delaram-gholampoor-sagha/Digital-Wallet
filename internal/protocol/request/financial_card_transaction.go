package request

type RegisterCardTransaction struct {
	TransactionGroupID int
	FinancialCardID    int
	Amount             float64
	Description        string
}

type Transfer struct {
	SenderCardID   int
	ReceiverCardID int
	Amount         float64
	Description    string
}
