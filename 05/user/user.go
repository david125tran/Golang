package User

import (
	"errors"
	"fmt"
	"time"
)

// Global Variables -------------------------------------
// Structs
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Helpers -------------------------------------
type Admin struct {
	Email    string
	Password string
	User
}

// A method with a receive, 'u', and a struct being received 'user'
// as an argument.
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email string, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "--/--/--",
			createdAt: time.Now(),
		},
	}
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Error: First name, last name, and birth date are required.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
