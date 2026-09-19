package main

import (
	"errors"
	"fmt"
	"strings"
)

// Global Variables -------------------------------------
var revenue float64
var expenses float64
var taxRate float64

func main() {
	fmt.Println(strings.Repeat("- ", 40))

	// User input
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// EBT = Revenue - Expenses
	ebt := revenue - expenses

	// Profit = EBT * (1 - Tax Rate)
	profit := ebt * (1 - taxRate)

	// Ratio = EBT / Profit
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

	// calculateFinancials()

}

func calculateFinancials() {
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

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Error: Invalid amount.")
	}

	return 0, nil
}
