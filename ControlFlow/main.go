package main

import "fmt"

// addTwoNumbers:
// Takes two integers as input parameters and returns their sum as an integer.
func addTwoNumbers(a int, b int) int {
	return a + b
}

// addTwoNumbersByUserInput:
// Reads two floating-point numbers from the user input,
// calculates their sum, and prints the result formatted to 2 decimal places.
func addTwoNumbersByUserInput() {

	var a, b float64

	fmt.Println("Enter the first number")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number")
	fmt.Scanln(&b)

	ans := a + b
	fmt.Printf("The sum of %.2f and %.2f is: %.2f\n", a, b, ans)
}

func main() {

	// Call addTwoNumbersByUserInput to get input from user and display sum
	addTwoNumbersByUserInput()

	// Call addTwoNumbers with fixed inputs 10 and 10, then print the result
	ans := addTwoNumbers(10, 10)
	fmt.Println(ans)

	// Read two floating-point numbers from user input,
	// calculate their sum, and print the result
	var a, b float64

	fmt.Println("Enter the first number")
	fmt.Scanln(&a)
	fmt.Println("Enter the second number")
	fmt.Scanln(&b)

	result := a + b
	fmt.Println(" the ans is : ", result)

	// Read an integer from user input and check if it is positive
	// Then print whether the number is even or odd
	var num int

	fmt.Println("Enter the number")
	fmt.Scanln(&num)

	if num <= 0 {
		fmt.Println("enter number gt ", num)
		return
	}
	if num%2 == 0 {
		fmt.Println("this is even number ", num)
	} else {
		fmt.Println("this is odd number ", num)
	}

}
