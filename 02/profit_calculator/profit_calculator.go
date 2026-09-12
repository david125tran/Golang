package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	// User input
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// EBT = Revenue - Expenses
	ebt := revenue - expenses

	// Profit = EBT * (1 - Tax Rate)
	profit := ebt * (1 - taxRate)

	// Ratio = EBT / Profit
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
