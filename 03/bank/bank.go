package main

import (
	"example.com/bank/file_operations"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"strings"
)

// Global Variables -------------------------------------
const accountBalanceFile string = "balance.txt"

// Main -------------------------------------
func main() {
	presentOptions()

	fmt.Println("Reach us 24/7 at:", randomdata.PhoneNumber())

	var accountBalance, err = file_operations.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println(err)
	}

	file_operations.WriteFloatToFile(accountBalance, accountBalanceFile)

	for {
		var choice int
		fmt.Println("")
		fmt.Println(strings.Repeat("- ", 40))
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// Control flow w/choice
		switch choice {
		// (1) - Check Balance
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		// (2) - Deposit
		case 2:
			// User Intake
			fmt.Println("How much do you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			// Edge Case - Deposit <= $0
			if depositAmount <= 0 {
				fmt.Println("Error: Invalid ammount. Deposit must be more than $0.00.")
				// Deposit
			} else {
				accountBalance += depositAmount
				fmt.Println("Your updated balance is:", accountBalance)
				file_operations.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		// (3) - Withdrawl
		case 3:
			// User Intake
			fmt.Println("How much do you want to withdrawal: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			// Edge Case - Withdrawl <= 0
			if withdrawalAmount <= 0 {
				fmt.Println("Error: Invalid ammount. Withdrawal must be more than $0.00.")
				// Edge Case - Insufficient funds
			} else if withdrawalAmount > accountBalance {
				fmt.Println("Error: Insufficient funds.")
				// Withdrawal
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Your updated balance is:", accountBalance)
				file_operations.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		// (4) - Exit
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for banking with us.")
			return
		// Default
		default:
			fmt.Println("Error: Invalid selection.")
		}
	}
}
