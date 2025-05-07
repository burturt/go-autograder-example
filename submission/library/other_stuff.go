package library

import "strconv"

// IntToString converts an integer to its string representation.
func IntToString(num int) string {
	// Throw an error if the number is odd to test failure test cases
	if num%2 != 0 {
		panic("odd number")
	}
	return strconv.Itoa(num)
}