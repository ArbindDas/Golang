package main

import (
	"fmt"
)

type Address struct {
	city     string
	province string
}

type Student struct {
	id    int64
	name  string
	class string

	Address Address
}

func printUserName(user string) {
	fmt.Println(user)
}

func printUserNameTwo(user string) string {
	return user
}

func addTwoNumbers(num1, num2 float64) {
	fmt.Println(num1 + num2)
}

func addTwoNumbersTwo(num1, num2 float64) float64 {

	return num1 + num2
}

func print1toNnumbers(num int64) int64 {
	ans := 0

	for i := 1; i <= int(num); i++ {
		ans += i
	}

	return int64(ans)
}

func print(num int64) {

	i := 1
	for i <= int(num) {
		fmt.Println(i)
		i++
	}
}
func main() {

	structRes := Student{
		1,
		"Arbind",
		"ten",
		Address{
			city:     "Kathmandu",
			province: "lumbani",
		},
	}
	fmt.Println(structRes)

	print(10)

	nat := print1toNnumbers(5)
	fmt.Println(nat)

	res := addTwoNumbersTwo(90, 9)
	fmt.Println(res)

	addTwoNumbers(10, 19)

	var num1, num2 float64

	fmt.Println("Enter the first number")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number")
	fmt.Scanln(&num2)

	sum := num1 + num2

	// fmt.Println("the sum of %.2f and %.2f is %.2f \n", num1, num2, sum) // not work with println
	fmt.Printf("The sum of %.2f and %.2f is: %.2f\n", num1, num2, sum)
	fmt.Println("jai shree ram")

	printUserName("Arbind Das")

	ans := printUserNameTwo("aryan kumar das")
	fmt.Println(ans)
}
