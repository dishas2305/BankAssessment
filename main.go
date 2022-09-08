package main

import (
	"BANKASSESMENT/types"
	"BANKASSESMENT/utils"
	"fmt"
)

func main() {
	var view string
	var number, key int

	fmt.Println("-------BANK---------")
	fmt.Println("-------Create accounts-------")

	fmt.Scan(&number)
	for i := 0; i < number; i++ {
		X := utils.CreateAccount()
		utils.Store(X)
		fmt.Println("-------Account Created Successfully-------\n", X)

	}
	fmt.Println("View Account List: Y/N")
	fmt.Scan(&view)
	if view == "Y" {
		fmt.Println(utils.AccountList)
	} else {
		fmt.Println("--------------------------------------------------")
	}

	for {
	here:
		fmt.Println("Enter the account number of account you want to perform operations on:")
		fmt.Scan(&key)
		_, ok := utils.AccountList[key]
		if ok {
			fmt.Println("Select operation:\n 1. Print Account Details\n 2. Print Current Balance\n 3. Deposit/Withdraw amount\n 4. Add New Account\n 5. View Account Statement\n ")
			var option int
			fmt.Scan(&option)
			switch option {
			case 1:
				utils.AccountList[key].AccDetails()
			case 2:
				fmt.Println(utils.AccountList[key].Balance)
			case 3:
				if utils.AccountList[key].AccountType == 1 {
					types.CasaAcc(*utils.AccountList[key])
				} else if utils.AccountList[key].AccountType == 2 {
					types.DepositAcc(*utils.AccountList[key])
				} else if utils.AccountList[key].AccountType == 1 {
					types.LoanAcc(*utils.AccountList[key])
				}
			case 4:
				utils.CreateAccount()
			case 5:
				utils.AccountList[key].ViewStatement()
			default:
				fmt.Println("Invalid option")
			}

		} else {
			fmt.Println("Enter valid account number")
			goto here
		}

	}

}
