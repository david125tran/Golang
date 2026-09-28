package main

import "fmt"

// ------------- main -------------
func main() {
	// ------------- numbers -------------
	numbers := []int{1, 2, 3}
	fmt.Println("\n---- numbers ----")
	fmt.Println("numbers:", numbers)

	// ------------- Anonymous Function -------------
	// Instead of defining a separate named function, we define the transformation
	// function directly where it is passed into transformNumbers.
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println("\n---- Anonymous Function ----")
	fmt.Println("transformed:", transformed)

	// ------------- Closures -------------
	fmt.Println("\n---- Closures ----")
	// createTransformer returns new transformation functions.
	// The returned functions "remember" the factor value that was passed in.
	double := createTransformer(2)
	triple := createTransformer(3)

	// Apply the dynamically created doubling function.
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println("doubledNumbers:", doubledNumbers)

	// Apply the dynamically created tripling function.
	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println("tripledNumbers:", tripledNumbers)
}

// ------------- Helpers -------------
// transformNumbers applies the provided transformation function to every value
// in the input slice and returns a new slice containing the transformed results.
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// createTransformer returns a new function that multiplies a number by factor.
// The returned function is a closure because it remembers and continues to use
// the factor value even after createTransformer has finished executing.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
