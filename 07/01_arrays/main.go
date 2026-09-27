package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// ----- prices array -----
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	// ----- featuredPrices array -----
	featuredPrices := prices[0:2]
	featuredPrices = append(featuredPrices, 5.99, 12.99, 1000.0)

	// ----- productNames array -----
	var productNames = [4]string{"A Book"}
	productNames[2] = "A Carpet" // Set value at index 2

	// ----- Working w/Arrays -----
	fmt.Println("\n----- prices array -----")
	fmt.Println(prices)
	fmt.Println("The 1st to 3rd values in the array:", prices[0:2])

	fmt.Println("The 1st value in the array:", prices[0]) // 10.99
	fmt.Println("The 2nd value in the array:", prices[1]) // 9.99
	fmt.Println("The 3rd value in the array:", prices[2]) // 45.99
	fmt.Println("The 4th value in the array:", prices[3]) // 20.0

	fmt.Println("\n----- featuredPrices array -----")
	fmt.Println(featuredPrices)
	fmt.Println("Length of featuredPrices:", len(featuredPrices))
	fmt.Println("Capacity of featuredPrices:", cap(featuredPrices))

	fmt.Println("\n----- productNames array -----")
	fmt.Println(productNames)
	fmt.Println("The 1st value in the array:", productNames[0]) // "A Book"
	fmt.Println("The 2nd value in the array:", productNames[1])
	fmt.Println("The 3rd value in the array:", productNames[2]) // "A Carpet"
	fmt.Println("The 4th value in the array:", productNames[3])
}
