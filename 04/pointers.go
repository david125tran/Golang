package main

import (
	"fmt"
)

func main() {
	// Regular variable
	age := 100

	fmt.Println("Age Value:", age)

	// No pointer
	adultYearsWithoutPointer := getAdultYearsWithoutPointer((age))
	fmt.Println("adultYearsWithoutPointer Value", adultYearsWithoutPointer)

	// Pointer Variable
	agePointerReference := &age
	fmt.Println("agePointerReference Address:", agePointerReference)
	fmt.Println("agePointerReference Value:", *agePointerReference) // Dereferencing
	agePointer := &age
	fmt.Println("agePointer:", getAdultYearsWithPointer(agePointer))

}

// Helper Functions -------------------------------------

// ------- Function without a Pointer -------
// getAdultYearsWithoutPointer returns the number of years the given age has
// been an adult. The age is passed by value, so the function receives a copy
// of the value.
func getAdultYearsWithoutPointer(age int) int {
	return age - 10
}

// ------- Function with a Pointer -------
// getAdultYears returns the number of years the given age has
// been an adult. The age is passed by pointer, so the function receives an
// address to age. The age is dereferenced so math can be performed.
func getAdultYearsWithPointer(age *int) int {
	return *age - 20 // Dereference pointer to perform subtraction

	// Alternatively, change the pointer's value which must also be dereferenced.
	// *age = *age - 18
}
