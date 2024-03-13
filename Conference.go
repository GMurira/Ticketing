package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Confernce"
	const noOfTickets = 50
	var bookings []string

	//Ask a user for their username.

	for {
		var firstName string
		var lastname string
		var email string
		var ticketsBought int
		fmt.Printf("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name")
		fmt.Scan(&lastname)
		fmt.Printf("Enter Your e-mail address:")
		fmt.Scan(&email)
		fmt.Printf("Enter number of tickets:")
		fmt.Scan(&ticketsBought)

		remainingTickets := 0
		remainingTickets = noOfTickets - ticketsBought

		bookings = append(bookings, firstName+""+lastname)

		fmt.Printf("The whole bookings slice %v\n", bookings)
		fmt.Printf("The first elemrnt of the slice %v\n", bookings[0])
		fmt.Printf("Slice type %T\n", bookings)
		fmt.Printf("Slice lengtn : %v\n", len(bookings))

		fmt.Printf("Welcome %v%v for booking a ticket number of tickets %v you will recieve a confirmation e-mail at %v\n", firstName, lastname, ticketsBought, email)
		fmt.Printf("This many tickets remainng %v for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are your bookings%v", bookings)
	}

}
