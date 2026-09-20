package main

import (
	"example.com/structs/user"
	"fmt"
)

// Global Variables -------------------------------------

// Main -------------------------------------
func main() {
	userFirstName := User.GetUserData("Please enter your first name: ")
	userLastName := User.GetUserData("Please enter your last name: ")
	userBirthDate := User.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *User.User

	appUser, error := User.New(userFirstName, userLastName, userBirthDate)

	if error != nil {
		fmt.Println(error)
		return
	}

	// fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)

	appUser.OutputUserDetails()

	appUser.ClearUsername()

	appUser.OutputUserDetails()

	// Admin
	admin := User.NewAdmin("davidtran@gmail.com", "123456")
	admin.User.OutputUserDetails()
}
