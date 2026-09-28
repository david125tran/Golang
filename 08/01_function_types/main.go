package main

import (
	"fmt"
)

// ------------- Type Aliases -------------

// transformFn defines the shape of a transformation function.
// Any function that accepts an int and returns an int can be used as a transformFn.
type transformFn func(int) int

// ------------- Main -------------

func main() {
	// Create the original slice of numbers.
	numbers := []int{1, 2, 3, 4}

	// Pass the double and triple functions into transformNumbers.
	// This demonstrates that functions can be passed as arguments in Go.
	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println("doubledNumbers:", doubledNumbers)
	fmt.Println("tripledNumber:", tripledNumbers)

	// Create another slice whose first value will determine
	// which transformation function should be returned.
	moreNumbers := []int{1, 1, 2}

	// getTransformerFunction examines moreNumbers.
	// Because the first value is 1, it returns the double function.
	transformerFn1 := getTransformerFunction(&moreNumbers)

	fmt.Println("moreNumbers:", moreNumbers)

	// Apply the dynamically selected transformation function.
	transformedMoreNumbers := transformNumbers(&moreNumbers, transformerFn1)
	fmt.Println("transformedMoreNumbers:", transformedMoreNumbers)

	// Check the newly transformed slice again.
	// Its first value is now 2, so getTransformerFunction returns triple.
	transformerFn2 := getTransformerFunction(&transformedMoreNumbers)

	// Reassign transformedMoreNumbers using the second transformation function.
	transformedMoreNumbers = transformNumbers(&transformedMoreNumbers, transformerFn2)
	fmt.Println("transformedMoreNumbers:", transformedMoreNumbers)
}

// ------------- Helpers -------------

// transformNumbers applies a transformation function to every integer in a slice.
// It returns a new slice containing the transformed values while leaving the
// original slice unchanged.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformedNumbers := []int{}

	for _, val := range *numbers {
		transformedNumbers = append(transformedNumbers, transform(val))
	}

	return transformedNumbers
}

// double returns the provided integer multiplied by 2.
func double(number int) int {
	return number * 2
}

// triple returns the provided integer multiplied by 3.
func triple(number int) int {
	return number * 3
}

// getTransformerFunction inspects the first value in the provided slice.
// If the first value is 1, it returns the double function; otherwise,
// it returns the triple function.
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}

	return triple
}
