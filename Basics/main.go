package main

import (
	"fmt"
	"math/rand"
)

// printUserName prints the given user name.
func printUserName(user string) {
	fmt.Println("The username is ->", user)
}

// getDefaultProductName returns a default product name as a string.
func getDefaultProductName() string {
	return "t-shirts"
}

// getProductName returns the provided product name.
func getProductName(product string) string {
	return product
}

// printUserId prints the given user ID.
func printUserId(id int) {
	fmt.Println("The user ID is:", id)
}

// getRandomUserId returns a randomly generated user ID between 0–999.
// In Go 1.20+, rand.Seed is not required for randomness.
func getRandomUserId() int {
	return rand.Intn(1000)
}

// getUserIdWithInput returns the same user ID that is passed as a parameter.
func getUserIdWithInput(id int) int {
	return id
}

func main() {
	// 1. Short variable declaration
	first := 100
	fmt.Println(first) // prints: 100

	// 2. Explicit var declaration with type
	var second int = 1000
	fmt.Println(second) // prints: 1000

	// 3. Var declaration with type inference
	var third = 900
	fmt.Println(third) // prints: 900

	printUserName("Arbind Das")

	product := getProductName("jeans")
	fmt.Println("Product name:", product)

	defaultProduct := getDefaultProductName()
	fmt.Println("Default product:", defaultProduct)

	printUserId(789)

	randomId := getRandomUserId()
	fmt.Println("Random User ID:", randomId)

	inputId := getUserIdWithInput(1000)
	fmt.Println("User ID from input:", inputId)

	var name string
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("hello ", name)

	var num1, num2 float64

	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Printf("The sum of %.2f and %.2f is: %.2f\n", num1, num2, sum)

}
