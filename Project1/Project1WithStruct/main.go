package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	TotalBalance int
}

// NewBankAccount creates a new bank account with an initial balance
func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{TotalBalance: initialBalance}
}

func (b *BankAccount) DepositeBalance(Balance int) {
	b.TotalBalance += Balance
	fmt.Println("Currently Deposite Amount Is ", Balance)
	fmt.Println("Total Balance Is ", b.TotalBalance)
}

func (b *BankAccount) WithDrawMoney(Balance int) error {

	if Balance > b.TotalBalance {
		return errors.New("withdrawal amount exceeds available balance")
	}

	b.TotalBalance -= Balance
	fmt.Println("Balance Withdraw ", Balance)
	fmt.Println("Remaning Balance ", b.TotalBalance)
	return nil
}

func (b *BankAccount) CheckBalance() {
	fmt.Println("Balance ", b.TotalBalance)
}

func main() {
	fmt.Println("jai shree ram")

	// ArbindAccount := BankAccount{1000}

	ArbindAccount := NewBankAccount(1000)
	ArbindAccount.DepositeBalance(4000)
	ArbindAccount.DepositeBalance(9000)
	if err := ArbindAccount.WithDrawMoney(2000); err != nil {
		fmt.Println("Error", err)
	}

	ArbindAccount.CheckBalance()
}
