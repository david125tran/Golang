package main

import (
	"fmt"
)

// ------------- Main -------------

func main() {
	// Call sumup with multiple integer arguments.
	// Because sumup is variadic, we can pass any number of ints directly.
	sum1 := sumup(1, 10, 15)
	fmt.Println("sum1:", sum1)
	sum2 := sumup(1, 10, 15, 30, 10)
	fmt.Println("sum2:", sum2)

	// The ... after numbers "unpacks" the slice so each element
	// is passed into sumup as an individual argument.
	numbers := []int{1, 30, 45}
	anotherSum := sumup(numbers...)
	fmt.Println("anotherSum:", anotherSum)

}

// ------------- Helpers -------------

// Non-variadic version:
// This version requires the caller to create and pass a slice of integers.
//
// func sumup(numbers []int) int {
// 	sum := 0
//
// 	for _, val := range numbers {
// 		sum = sum + val
// 	}
//
// 	return sum
// }

// sumup accepts any number of integer arguments and returns their total.
// The ...int syntax makes numbers a variadic parameter, which Go treats
// as a slice of ints inside the function.
func sumup(numbers ...int) int {
	sum := 0

	// Loop through every number passed into the function
	// and add each value to the running total.
	for _, val := range numbers {
		sum = sum + val
	}

	return sum
}
