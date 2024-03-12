package main

import "fmt"

func main() {
	conferenceName := "Go Confernce"
	const noOfTickets = 50

	//Ask a user for their username.

	//Array for booking
	var bookings [50]string
	var firstName string
	var email string
	var ticketsBought int
	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Printf("Enter Your e-mail address:")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&ticketsBought)

	remainingTickets := 0
	remainingTickets = noOfTickets - ticketsBought

	bookings[0] = firstName

	fmt.Printf("The whole bookings array %v\n", bookings)
	fmt.Printf("The first elemrnt of the array %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array lengtn : %v\n", len(bookings))

	fmt.Printf("Welcome %v for booking a ticket number of tickets %v you will recieve a confirmation e-mail at %v\n", firstName, ticketsBought, email)
	fmt.Printf("This many tickets remainng %v for %v\n", remainingTickets, conferenceName)
}
