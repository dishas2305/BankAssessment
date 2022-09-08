package types

type AccountOperations interface {
	Deposit()
	Withdrawal()
}

type CaSa interface {
	AccountOperations
}

type Deposit interface {
	Deposit()
}

type Loan interface {
	Deposit()
}
