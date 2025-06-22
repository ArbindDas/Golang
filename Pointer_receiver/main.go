package main

import "fmt"

type Calculator struct{}

// Method with value receiver
func (c Calculator) AddValue(a int, b int) {
	fmt.Println("Value receiver:", a+b)
}

// Method with pointer receiver
func (c *Calculator) AddPointer(a int, b int) {
	fmt.Println("Pointer receiver:", a+b)
}

func main() {
	fmt.Println("jai shree ram")

	c := Calculator{}
	c.AddValue(10, 10) // call method with value receiver

	cp := &Calculator{}
	cp.AddPointer(10, 10) // call method with pointer receiver
}
