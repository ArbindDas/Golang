package main

import "fmt"

func DepositeBalance(balance int, totalBalance *int) {

	*totalBalance += balance
	fmt.Println("Total Balance is : ", *totalBalance)
}

func WithDrawMoney(balance int, totalBalance *int) {
	if balance > *totalBalance {
		fmt.Println("total balance is less then your withdraw amout")
		return
	}
	if balance <= *totalBalance {
		*totalBalance -= balance
		fmt.Println("remaing balance of your account is ", *totalBalance)

	}
}

func CheckBalance(total *int) {
	fmt.Println("Current Balance : ", *total)
}
func main() {

	total := 0

	DepositeBalance(10, &total)
	DepositeBalance(90, &total)

	WithDrawMoney(80, &total)

	CheckBalance(&total)

}
