package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Global Variables -------------------------------------
const accountBalanceFile = "balance.txt"

// Helper Functions -------------------------------------
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	// File not found
	if err != nil {
		return 1000, errors.New("Error: Failed to read file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	// Conversion error
	if err != nil {
		return 1000, errors.New("Error: Failed to parse stored balance value.")
	}

	return balance, nil
}

// Main -------------------------------------
func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println(err)
	}

	writeBalanceToFile(accountBalance)

	for {
		fmt.Println(strings.Repeat("- ", 40))

		fmt.Println("Welcome to the Bank!")
		fmt.Println(("What do you want to do?"))
		fmt.Println("1.	Check balance.")
		fmt.Println("2.	Deposit money.")
		fmt.Println("3.	Withdraw money.")
		fmt.Println("4.	Exit..")

		var choice int

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
				writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
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

		// // Control flow w/if, else
		// // (1) - Check Balance
		// if choice == 1 {
		// 	fmt.Println("Your balance is:", accountBalance)
		// 	// (2) - Deposit
		// } else if choice == 2 {
		// 	// User Intake
		// 	fmt.Println("How much do you want to deposit: ")
		// 	var depositAmount float64
		// 	fmt.Scan(&depositAmount)

		// 	// Edge Case - Deposit <= $0
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Error: Invalid ammount. Deposit must be more than $0.00.")
		// 		// Deposit
		// 	} else {
		// 		accountBalance += depositAmount
		// 		fmt.Println("Your updated balance is:", accountBalance)
		// 	}
		// 	// (3) - Withdrawl
		// } else if choice == 3 {
		// 	// User Intake
		// 	fmt.Println("How much do you want to withdrawal: ")
		// 	var withdrawalAmount float64
		// 	fmt.Scan(&withdrawalAmount)

		// 	// Edge Case - Withdrawl <= 0
		// 	if withdrawalAmount <= 0 {
		// 		fmt.Println("Error: Invalid ammount. Withdrawal must be more than $0.00.")
		// 		// Edge Case - Insufficient funds
		// 	} else if withdrawalAmount > accountBalance {
		// 		fmt.Println("Error: Insufficient funds.")
		// 		// Withdrawal
		// 	} else {
		// 		accountBalance -= withdrawalAmount
		// 		fmt.Println("Your updated balance is:", accountBalance)
		// 	}
		// 	// (4) - Exit
		// } else {
		// 				fmt.Println("Goodbye!")c
		// 				breakc
		// }
	}

	// fmt.Println("Thanks for banking with us.")
}
