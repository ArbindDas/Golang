package main

import "fmt"

type Addition struct {
	value1 float64
	value2 float64
}

func initializeValues(num1, num2 float64) *Addition {

	return &Addition{value1: num1, value2: num2}
}

func (a *Addition) addTwo() {
	ans := a.value1 + a.value2
	fmt.Println(ans)
}
func addTwoNumberWithUsingPointers(num1, num2 *int) {

	ans := *num1 + *num2

	fmt.Printf("the addition of %d + %d is  = %d\n ", *num1, *num2, ans)
}
func main() {
	// & gives you the Memory address of variable (a pointer)
	// * Takes a pointer (memory address) and gives you the actual value stored at the address (dereference the pointer)

	number1 := 100
	number2 := 100

	addTwoNumberWithUsingPointers(&number1, &number2)

	fmt.Println("Jai Shree Ram")

	ans := initializeValues(10, 10)
	ans.addTwo()
	fmt.Println(ans)

	result := Addition{
		value1: 100,

		value2: 100,
	}

	fmt.Println(result)
}
