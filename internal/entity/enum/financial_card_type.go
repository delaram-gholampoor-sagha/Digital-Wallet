package enum


type FinancialCardType uint

const (
    Debit  FinancialCardType = iota
    Credit         
    Gift            
)
