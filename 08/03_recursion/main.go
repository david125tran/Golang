package main

import (
	"fmt"
)

// ------------- Main -------------

func main() {
	// Calculate the factorial of 5.
	// 5! = 5 × 4 × 3 × 2 × 1 = 120
	fact := factorial(5)

	fmt.Println("fact:", fact)
}

// ------------- Helpers -------------

// Iterative version:
// This version calculates the factorial using a loop.
//
// func factorial(number int) int {
// 	result := 1
//
// 	for i := 1; i <= number; i++ {
// 		result = result * i
// 	}
//
// 	return result
// }

// factorial returns the factorial of the provided integer using recursion.
// The function keeps calling itself with number - 1 until it reaches 0,
// which acts as the base case and stops the recursion.
func factorial(number int) int {
	// Base case:
	// 0! is defined as 1, and this prevents infinite recursive calls.
	if number == 0 {
		return 1
	}

	// Recursive case:
	// Multiply the current number by the factorial of the next smaller number.
	return number * factorial(number-1)
}
