package main

import (
	"fmt"
)

// Type Aliases
type floatMap map[string]float64

func main() {
	// ----- Maps -----
	websites := map[string]string{
		"AWS":      "https://www.aws.com",
		"Google":   "https://google.com",
		"Facebook": "https://facebook.com",
	}
	// Add key:value pair
	websites["LinkedIn"] = "https://linkedin.com"
	// Delete key:value pair
	delete(websites, "Facebook")

	// ----- Working w/Maps -----
	fmt.Println("\n----- websites Map -----")
	fmt.Println("websites map:\n", websites)
	fmt.Println("AWS value:", websites["AWS"])

	fmt.Println("\n----- Looping Through websites Map -----")
	for index, value := range websites {
		fmt.Println("Index:", index, "\nValue:", value)
	}

	// ----- Special 'make' Function (arrays) -----
	// make creates a slice with a defined length and optional capacity.
	// Preallocating capacity can improve performance by reducing how often Go needs
	// to allocate a larger backing array as new elements are appended.

	userNames := make([]string, 2)

	userNames[0] = "Hannah"
	userNames[1] = "Banana"

	fmt.Println("\n----- userNames ('make' function) -----")
	fmt.Println("userNames", userNames)

	fmt.Println("\n----- Looping Through userNames Array -----")
	for index, value := range userNames {
		fmt.Println("Index:", index, "\nValue:", value)
	}

	// ----- Special 'make' Function (maps) -----
	// courseRatings := make(map[string]float64, 2)

	// Making the map with a type alias
	courseRatings := make(floatMap, 2)

	courseRatings["go"] = 4.9
	courseRatings["react"] = 4.1

	fmt.Println("\n----- courseRatings ('make' function) -----")
	fmt.Println("courseRatings:\n", courseRatings)
}
