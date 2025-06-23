package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	TotalBalance float64

	//  Private fields	Unexported (lowercase) fields
	// Public fields	Exported (uppercase) fields
}

func CreateUserAccountWithMoney(initialBalance float64) *BankAccount {
	return &BankAccount{TotalBalance: initialBalance}
}

func (b *BankAccount) depositeBalance(balance float64) {
	fmt.Println("Before deposited the actual balance is", b.TotalBalance)
	b.TotalBalance += balance
	fmt.Println("total balance is ", b.TotalBalance)
}

func (b *BankAccount) withDrawMoney(balance float64) error {

	if balance > b.TotalBalance {

		return errors.New("Withdraw amount execeeds availibale balance")
	}

	b.TotalBalance -= balance
	fmt.Println("successfully withdraw money .....", balance)
	fmt.Println("Availibale balance is ...", b.TotalBalance)
	return nil

}

func (b *BankAccount) checkBalance() {
	fmt.Println("Total balance is ", b.TotalBalance)
}
func main() {

	Arbind := CreateUserAccountWithMoney(1000)

	Arbind.depositeBalance(1000)
	Arbind.depositeBalance(1000)

	if err := Arbind.withDrawMoney(4000); err != nil {
		fmt.Println("error", err)
	}

	Arbind.checkBalance()

	fmt.Println("jai shree ram")
}
