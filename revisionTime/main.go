package main

import "fmt"

// AddValues holds two float64 numbers to perform addition
type AddValues struct {
	value1 float64
	value2 float64
}

// NewAddValues is a constructor that initializes and returns a pointer to AddValues
func NewAddValues(num1, num2 float64) *AddValues {
	return &AddValues{value1: num1, value2: num2}
}

// Sum returns the addition of the two float64 values
// public function
func (a *AddValues) Sum() float64 {
	return a.value1 + a.value2
}

// addIntPointers adds two integers via pointers and returns the result
// private function
func addIntPointers(num1, num2 *int) int {
	return *num1 + *num2
}

func main() {
	num1 := 100
	num2 := 100

	result := addIntPointers(&num1, &num2)
	fmt.Printf("The addition of %d + %d = %d\n", num1, num2, result)

	fmt.Println("Jai Shree Ram")

	addVals := NewAddValues(10.0, 10.0)
	sum := addVals.Sum()
	fmt.Printf("The sum of %.2f and %.2f = %.2f\n", addVals.value1, addVals.value2, sum)
}
