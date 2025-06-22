package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Age     int
	Address // Embedded type
}

func main() {

	e := Employee{
		Name: "Arbind",
		Age:  21,
		Address: Address{
			City:    "jeetpur",
			Country: "Nepal",
		},
	}

	fmt.Println("Name -> ", e.Name)
	fmt.Println("age -> ", e.Age)
	fmt.Println("city ->", e.City) // Access directly via embedding
	fmt.Println("Country ->", e.Country)

}
