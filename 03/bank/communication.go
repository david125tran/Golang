package main

import (
	"fmt"
	"strings"
)

// Helper Functions -------------------------------------
func presentOptions() {
	fmt.Println(strings.Repeat("- ", 40))

	fmt.Println("Welcome to the Bank!")
	fmt.Println(("What do you want to do?"))
	fmt.Println("1.	Check balance.")
	fmt.Println("2.	Deposit money.")
	fmt.Println("3.	Withdraw money.")
	fmt.Println("4.	Exit..")

}
