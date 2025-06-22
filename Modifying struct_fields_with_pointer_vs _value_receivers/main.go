package main

import "fmt"

type Calculator struct {
	total int
}

// Method with **value receiver** - works on a copy of Calculator
func (c Calculator) AddValue(a int) {
	c.total += a
	fmt.Println("Inside AddValue:", c.total)
}

// Method with **pointer receiver** - works on the original Calculator
func (c *Calculator) AddPointer(a int) {
	c.total += a
	fmt.Println("Inside AddPointer:", c.total)
}

func main() {
	c := Calculator{} // Create a Calculator instance

	c.AddValue(10)
	fmt.Println("After AddValue call:", c.total) // total won't change outside AddValue

	c.AddPointer(10)
	fmt.Println("After AddPointer call:", c.total) // total will be updated

	// Also possible to call pointer receiver method on value, Go takes address automatically
	c2 := Calculator{}
	c2.AddPointer(5)
	fmt.Println("After AddPointer call on c2:", c2.total)
}
