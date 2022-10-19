package main

import (
	"fmt"
	"strings"
)

const conferenceName string = "Go Conference"

var conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInputs()

		isValidInputs := validateUserInputs(firstName, lastName, email, userTickets)

		if isValidInputs {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()

			fmt.Printf("These are all the first names: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("There no more tickets. Come back next year.")
				break
			}

		} else {
			fmt.Printf("There is only %v left. So you can't book %v \n", remainingTickets, userTickets)
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName && isValidEmail && isValidTickets
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name (pointer)
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets! You will be receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v is remaining for %v\n", remainingTickets, conferenceName)
}
