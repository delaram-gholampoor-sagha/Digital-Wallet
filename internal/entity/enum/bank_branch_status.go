package enum



type BankBranchStatus uint

const (
	BankBranchInactive BankBranchStatus = iota
	BankBranchActive
)