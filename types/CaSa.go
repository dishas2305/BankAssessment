package types

import (
	"BANKASSESMENT/utils"
	"fmt"
)

func CasaAcc(a utils.Account) {
	var acc CaSa = a
	var option int
	fmt.Print("1.Deposit amount\n2.Withdraw amount\n")
	fmt.Scan(&option)
	if option == 1 {
		acc.Deposit()

	} else if option == 2 {
		acc.Withdrawal()
	} else {
		fmt.Println("Invalid entry")
	}
}
