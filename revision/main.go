package main

import "fmt"

// add two number using function
// without return anything void function
func addTwoNumber(num1, num2 int) {
	ans := num1 + num2
	fmt.Printf("the addition of %d and %d is : %d\n", num1, num2, ans)
}

// add two number with return value
func addTwoNumberWithReturns(num1, num2 int) int {

	return num1 + num2
}

func addTwoNumberWithPointer(num1, num2 *int) {

	ans := *num1 + *num2
	fmt.Println(ans)
}

func addTwoNumberWithPointerWithReturns(num1, num2 *int) int {

	return *num1 + *num2
}

func printProductsData(products []int) {

	fmt.Println("With index and both values")
	for i, val := range products {
		fmt.Printf("Index %d , values %d\n", i, val)
	}

	fmt.Println("with values only")

	for _, values := range products {
		fmt.Printf("values %d\n", values)
	}
}
func main() {
	fmt.Println("jai shree ram")

	addTwoNumber(10, 99)

	fmt.Println(addTwoNumberWithReturns(90, 90))

	num1 := 100
	num2 := 100

	addTwoNumberWithPointer(&num1, &num2)

	fmt.Println(addTwoNumberWithPointerWithReturns(&num1, &num2))

	//The & operator means: get the memory address of the struct.
	// * works with the value at that address.

	numbers := []int{10, 10, 10, 20}

	// fmt.Println(numbers)

	// for i, val := range numbers {
	// 	fmt.Printf("index %d  values %d\n", i, val)
	// }

	printProductsData(numbers)

	ans := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
	}

	fmt.Println(ans)
	// Accessing a value
	fmt.Printf("the age of ALice is : %d\n", ans["Alice"])

	// Adding a new key-value pair
	ans["Aryan"] = 32

	// Updating a value
	ans["Alice"] = 90

	// Deleting a key
	delete(ans, "Bob")

	fmt.Println(ans)

	// Looping through the map

	for name, age := range ans {
		fmt.Printf("%s is %d years old\n", name, age)

	}

	// Check if a key exists
	age, exists := ans["Alice"]

	if exists {
		fmt.Printf("Alice age  is %d\n", age)
	} else {
		fmt.Printf("ALice is not present in the map")
	}

}
