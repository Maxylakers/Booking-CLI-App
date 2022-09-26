package main

import (
	"fmt"
)

func main() {
	conferenceName := "Max Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

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

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}