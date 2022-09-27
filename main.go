package main

import (
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Max Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {
		// Call function getUserInput
		firstName, lastName, email, userTickets := getUserInput()

		// Call function to validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Call function bookTicket
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			// Call Function Print first name
			firstNames := getFirstName()
			fmt.Printf("The first names of all bookings are:  %v\n", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Printf("Our tickets are booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The name you have entered is too short. Try again")
			}
			if !isValidEmail {
				fmt.Println("You have entered an invalid email. Try again")
			}
			if !isValidTicketNumber {
				fmt.Println("You have entered an invalid number of tickets.")
			}

		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome To %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Here Now To Attend")
}

func getFirstName() []string {
	// To use the for each statement, we need the "index" and the element declared. But since we are not using the index variable in the program, we will replace it with an "_"
	firstNames := []string{}
	for _, booking := range bookings {
		// We will use the strings.Fields builtin function to seperate each element in booking with whitespaces
		// names := strings.Fields(booking)
		//Now we are changing the initial to a Map
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	var remainingTickets = remainingTickets - userTickets

	// Replacing the Map block below with struct

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//Create map for user data

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickest"] = strconv.FormatUint(uint64(userTickets), 10)

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("The List of bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We now have: %v tickets remaining for the %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets to %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \n via email address %v\n", ticket, email)
	fmt.Println("##########")
}
