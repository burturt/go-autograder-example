package main

import (
	"fmt"
	"math/rand"
)

// AddTwoInts adds two integers and returns the result.
func AddTwoInts(a int, b int) int {
	return a + b
}

func AddTwoIntsUnreliable(a int, b int) int {
	// Randomly returns wrong value 25% of the time for testing retesting
	if rand.Intn(4) == 0 {
		return a + b + 2
	}
	return a + b
}

func main() {
	// Example usage
	result := AddTwoInts(3, 5)
	fmt.Println("The sum is:", result)
}