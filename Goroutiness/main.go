package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) // Add one goroutine to wait for
		go func(n int) {
			defer wg.Done() // Mark this goroutine as done
			fmt.Println(n)
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish
}
