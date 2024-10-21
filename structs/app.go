package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct { // struct = data group
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}

/**
* Constructor
 */
func newUser(firstName string, lastName string, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("firstname, lastname and birthdate are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		age:       19,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Enter your firstname: ")
	userLastName := getUserData("Enter your lastname: ")
	userBirthDate := getUserData("Enter your birthdate (MM/DD/YYY): ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}

func (u user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUsername() { // without star we will edited a copy record
	u.firstName = ""
	u.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
