package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Max Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome To %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Here Now To Attend")

	for {
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

		if userTickets <= remainingTickets {
			var remainingTickets uint = remainingTickets - userTickets

			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We now have: %v tickets remaining for the %v\n", remainingTickets, conferenceName)

			// To use the for each statement, we need the "index" and the element declared. But since we are not using the index variable in the program, we will replace it with an "_"
			firstNames := []string{}
			for _, booking := range bookings {
				// We will use the strings.Fields builtin function to seperate each element in booking with whitespaces
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of all bookings are:  %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end the program
				fmt.Printf("Our tickets are booked out. Come back next year")
				break
			}
		} else {

			fmt.Printf("We currently have only %v number of tickets remaining, you cannot book %v tickets\n", remainingTickets, userTickets)
		}

	}

}
