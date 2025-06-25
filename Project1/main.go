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

	CheckBalance(&total) //The & operator means: get the memory address of the struct.

	// 	&total sends the address of total to the function.

	// totalBalance *int receives that address.

	// *totalBalance works with the value at that address.

}
