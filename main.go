package main

import (
	"fmt"
)

func main() {
	conferenceName := "Max Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome To %v Booking Application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Here Now To Attend")

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

	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Content of the whole array: %v\n", bookings)
	fmt.Printf("First value of the array: %v\n", bookings[0])
	fmt.Printf("Type of the array: %T\n", bookings)
	fmt.Printf("Length of the array: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We still have: %v tickets remaining for the %v", remainingTickets, conferenceName)

}
