package types

import (
	"BANKASSESMENT/utils"
	"fmt"
)

func DepositAcc(a utils.Account) {
	var acc Deposit = a
	if a.AccountSubType == 1 {
		acc.Deposit()
	} else if a.AccountSubType == 2 {
		acc.Deposit()
	} else {
		fmt.Println("Invalid entry")
	}

}
