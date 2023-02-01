// Define the package which helper belongs to --> main
package main

import "strings"

// Function to validate user input
// Input values in the first parenthesis,
// and list of output[return] params outside braces

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}
