package main

import (
	"fmt"
	"math"
)

// Global Variables -------------------------------------
// Constant - Fixed, immutable value
const inflationRate float64 = 2.5

func main() {
	// Variables & Data Types -------------------------------------
	// Type - Explicit declaration
	// var investmentAmount float64 = 1000
	// var years float64 = 10

	// Type - Inferring with var
	// var expectedReturnRate = 5.5

	// Type - Inferring with `:=`
	// expectedReturnRate := 5.5

	// Type - float64 explicit declaration by adding decimals
	// var years = 10.0

	// Declaring multiple variables
	// var investmentAmount, years float64 = 1000, 10

	// Declaring multiple variables with float64 explicit declaration by adding decimals
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	// Variables -------------------------------------
	investmentAmount := 1000.0
	years := 10.0
	expectedReturnRate := 5.5

	// Constant - Fixed, immutable value
	// const inflationRate float64 = 2.5

	// Main -------------------------------------
	// Fetch user input
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)

	// Calculations
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Print w/string concatenation
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (Inflation Adjusted):", futureRealValue)

	// Print w/formatted string literal
	// fmt.Printf("Future Value: %v\n", futureValue)
	// fmt.Printf("Future Value (Inflation Adjusted): %v\n", futureRealValue)

	// fmt.Printf("Future Value: %.1f\n", futureValue)
	// fmt.Printf("Future Value (Inflation Adjusted): %.1f\n", futureRealValue)

	// String return
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Inflation Adjusted): %.1f\n", futureRealValue)
	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

	// Building multiline strings
	// fmt.Printf(`Future Value: %.1f
	// Future Value (Inflation Adjusted): %.1f`,
	// 	futureValue, futureRealValue)

	// Calling custom function
	outputText("Hello World")

	futureValue, futureRealValue = calculateFutureValues(investmentAmount, expectedReturnRate, years)

}

// Functions -------------------------------------
func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(
	investmentAmount float64,
	expectedReturnRate float64,
	years float64,
) (float64, float64) {

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue

}

// Alternative function declaration w/type declaration
// func calculateFutureValues(
// 	investmentAmount float64,
// 	expectedReturnRate float64,
// 	years float64,
// ) (futureValue float64,
// 	realFutureValue float64) {

// 	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	realFutureValue = futureValue / math.Pow(1+inflationRate/100, years)
// 	return

// }
