package main

import (
	"fmt"
)

func main() {

	fmt.Println("jai shree ram")

	numbersSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("orignial", numbersSlice, "Len", len(numbersSlice), "cap", cap(numbersSlice))
	numbersSlice = append(numbersSlice, 10, 20)
	// fmt.Println("Extended", extended, "Len", len(extended), "cap", cap(extended))

	fmt.Println(numbersSlice)
	for i, val := range numbersSlice {
		fmt.Printf("index:%d & value %d\n", i, val)
	}

	for _, v := range numbersSlice {
		fmt.Println(v)
	}

}
