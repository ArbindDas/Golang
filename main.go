// package main

// import (
// 	"fmt"
// )

// type SalarySet struct {
// 	Salary int
// }

// func CreateNewSalarySet(initialSalary int) *SalarySet {
// 	return &SalarySet{Salary: initialSalary}
// }

// func (s *SalarySet) printSalry() {

// 	tax := (s.Salary * 13) / 100

// 	tds := (s.Salary * 4) / 100

// 	total := (s.Salary - tax - tds)

// 	fmt.Println("tax ", tax)
// 	fmt.Println("tds ", tds)

// 	fmt.Println("total ", total)
// }

// func main() {

// 	sal := CreateNewSalarySet(1000)

// 	sal.printSalry()

// 	// 	& operator: Gives you the memory address of a variable (a pointer).

// 	// * operator: Takes a pointer (memory address) and gives you the actual value stored at that address (dereferences the pointer).
// }

package main

import (
	"encoding/json"
	"fmt"
)

func addTwoNumber(num1, num2 int) {

	ans := num1 + num2

	fmt.Printf("the additon of %d & %d is  %d", num1, num2, ans)
	fmt.Println()

}

func addTwoNumberTwo(num1, num2 int) {
	addTwoNumber(num1, num2)
}

type Employee struct {
	Name    string
	Address string
	Number  string
}

func main() {

	// addTwoNumber(10, 10)

	// addTwoNumberTwo(90, 9)

	Arbind := Employee{
		Name:    "Arbind",
		Address: "jeetpur",
		Number:  "9821161214",
	}

	data, _ := json.MarshalIndent(Arbind, "", " ")
	fmt.Println(string(data))

	fmt.Println(Arbind)

	fmt.Println("jai shree ram")

	numbers := []int{1, 2, 90, 4}

	fmt.Println(numbers)

	for i, val := range numbers {
		fmt.Printf("index %d values %d\n", i, val)
	}

	for _, value := range numbers {
		fmt.Printf("values %d\n", value)
	}

	fmt.Println()
	products := []string{"jeans", "shirt", "vest", "pant"}

	for i, val := range products {
		fmt.Printf("index %d val %s\n", i, val)
	}

	for _, val := range products {
		// _ is used to ignore the index
		// val holds the actual value from the slice (like "jeans", "shirt", etc.)
		fmt.Printf("%s\n", val)
	}

}
