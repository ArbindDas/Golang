package main

import "fmt"

// addTwoNumbers takes two integers and prints their sum.
func addTwoNumbers(a int, b int) {
	fmt.Println(a + b)
}

// print takes a string (user name) and returns it.
func print(user string) string {
	return user
}

// sum1toNnum takes an integer and returns the sum of numbers from 1 to that integer.
func sum1toNnum(num int) int {
	var ans int
	for i := 1; i <= num; i++ {
		ans += i
	}
	return ans
}

func main() {
	fmt.Println("jai shree ram") // Greeting message

	addTwoNumbers(10, 10) // Calls the addTwoNumbers function to print 20

	ans := print("Arbind Das")           // Calls the print function and stores the returned name
	fmt.Println("the username is ", ans) // Prints the returned name

	result := sum1toNnum(5) // Calls the sum1toNnum function to get the sum 1+2+3+4+5 = 15
	fmt.Println(result)     // Prints the result
}
