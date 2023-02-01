package main

import (
	"fmt"
	"sync"
	"time"
)

// Package level variables
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName        string
	lastName         string
	email            string
	NumberOfTictkets uint
}

var wg = sync.WaitGroup{}

func main() {
	// Greet users
	greetUsers()

	// Iterate to get bookings sample:
	// for {
	// validate user Input
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

	// Check if userTickets is greater than remainingTickets
	if isValidName && isValidEmail && isValidTickets {

		// Returned function for booking ticket
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Function to print first name of users
		// return value saved into a variable: firstNames
		firstNames := getFirstName()
		fmt.Printf("These first names of bookings are %v\n", firstNames)

		// Check if ticket is equal to zero
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year!")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email address you entered doesn't contain the @ sign")
		}
		if !isValidTickets {
			fmt.Println("The number of tickets you entered is invalid")
		}
	}
	// }
	wg.Wait()
}

// Function to greet Users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

// Fucntion to get User's first name
func getFirstName() []string {
	// looping through firstNames
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	// Printing all firstNames
	return firstNames
}

// Function to get User Input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask for user's name
	// The pointer {&} is used here to get the
	// memory address of the firstName variable
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Function to book Ticket
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Users' data
	var userData = UserData{
		firstName:        firstName,
		lastName:         lastName,
		email:            email,
		NumberOfTictkets: userTickets,
	}

	// Bookings list with user information
	// as key value pairs
	bookings = append(bookings, userData)
	// Print the bookings list
	fmt.Printf("List of bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive an email confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets left for the %v\n", remainingTickets, conferenceName)
}

// Generate booked ticket
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Set delay 10 secs
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %vv %v", userTickets, firstName, lastName)
	fmt.Println("****************")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("****************")
	wg.Done()
}
