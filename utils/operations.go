package utils

import (
	"fmt"
	"time"
	//"strconv"
)

var DepositAmount, Amount float64

const INTEREST_RATE = 0.0350

type Account struct {
	AccountNumber  int
	AccountHolder  string
	InitialAmount  float64
	Balance        float64
	EmailID        string
	Mobile         string
	AccountType    int
	AccountSubType int
}

var AccountList = map[int]*Account{}

func Store(account Account) {
	AccountList[account.AccountNumber] = &account
}

func CreateAccount() (a Account) {
	fmt.Println("Enter Account No(13 digits):")
	fmt.Scanln(&a.AccountNumber)

	fmt.Println("Enter Account holder name:")
	fmt.Scanln(&a.AccountHolder)

	fmt.Println("Enter Initial Amount:")
	fmt.Scanln(&a.InitialAmount)

	a.Balance = a.InitialAmount

	fmt.Println("Enter Email ID:")
	fmt.Scanln(&a.EmailID)

	fmt.Println("Enter Mobile number:")
	fmt.Scanln(&a.Mobile)

	fmt.Println("Enter Account type: 1. Casa 2. Deposit 3. Loan")
	fmt.Scanln(&a.AccountType)

here:
	if a.AccountType == 1 {
		fmt.Println("Enter Account Subtype: 1. Current Account 2. Savings Account")
		fmt.Scanln(&a.AccountSubType)
	} else if a.AccountType == 2 {
		fmt.Println("Enter Account Subtype: 1. Fixed Deposit 2. Recuuring Deposit")
		fmt.Scanln(&a.AccountSubType)
	} else if a.AccountType == 3 {
		fmt.Println("Enter Account Subtype: 1. Housing loan 2. Vehilcle loan 3. Personal Loan")
		fmt.Scanln(&a.AccountSubType)
	} else {
		fmt.Println("Enter valid Account Type")
		goto here
	}

	return a
}

func (a Account) AccDetails() {
	fmt.Println(AccountList)
	fmt.Println("Account Number: ", a.AccountNumber)
	fmt.Println("Account Holder Name: ", a.AccountHolder)
	fmt.Println("Balace: ", a.Balance)
	fmt.Println("Email ID: ", a.EmailID)
	fmt.Println("Mobile: ", a.Mobile)
}

func (a Account) Deposit() {
	fmt.Printf("Enter amount you want to Deposit\n")
	fmt.Scan(&DepositAmount)
	a.Balance = a.InitialAmount + DepositAmount
	fmt.Println("Credited Amount: ", DepositAmount)
	fmt.Println("New Balance", a.Balance)
}

func (a Account) Withdrawal() {
	fmt.Printf("Enter the amount you want to withdraw\n")
	fmt.Scan(&Amount)
	if a.InitialAmount >= Amount || a.Balance >= Amount {
		a.Balance -= Amount
		fmt.Println("Amount Debited: ", Amount)
		fmt.Println("Balance after withdrawal", a.Balance)
	} else {
		fmt.Println("Your Balance is less than", Amount)
	}
}

func (a Account) ViewStatement() {
	fmt.Println("------------BANK------------")
	fmt.Println("------------Account Statement------------")
	fmt.Println("Date:", (time.Now().Day()), "Time:", time.Now())
	fmt.Println("Interest: ", a.ComputeAccuringInterest())
	fmt.Println("Balance: ", a.Balance)
}

func (a Account) ComputeAccuringInterest() (final_amount float64) {
	FinalAmount := a.Balance * (a.Balance * INTEREST_RATE / 365)
	return FinalAmount
}
