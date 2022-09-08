package types

import (
	"BANKASSESMENT/utils"
)

func LoanAcc(a utils.Account) {
	var acc Loan = a
	acc.Deposit()
}
