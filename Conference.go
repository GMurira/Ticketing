package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	welcome := "Welcome to the Annual"
	const noOfTickets = 50
	var bookings []userData

	remainingTickets := noOfTickets

	firstName, lastName, email, ticketsBought := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, ticketsBought, remainingTickets)

	wg.Add(1)
	go sendTicket(ticketsBought, firstName, lastName, email)

	if isValidName && isValidTicketNumber && isValidEmail {

		userData := userData{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: uint(ticketsBought),
		}

		bookings = append(bookings, userData)

		fmt.Printf("Welcome %v %v to the %s, thank you for booking %d tickets. You will receive a confirmation e-mail at %s\n", firstName, lastName, welcome, ticketsBought, email)

		firstNames := getFirstNames(bookings)
		fmt.Printf("The first Names of the bookings are %v\n", firstNames)

		fmt.Printf("These are your bookings: %v\n", bookings)

		remainingTickets -= ticketsBought
		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Printf("Our Conference is booked out, come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("The name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("E-mail address you entered does not contain @")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func getFirstNames(bookings []userData) []string {
	firstNames := make([]string, len(bookings))
	for i, booking := range bookings {
		firstNames[i] = booking.firstName
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var ticketsBought int

	fmt.Printf("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Printf("Enter Your e-mail address:")
	fmt.Scan(&email)
	fmt.Printf("Enter number of tickets:")
	fmt.Scan(&ticketsBought)

	return firstName, lastName, email, ticketsBought
}

func validateUserInput(firstName string, lastName string, email string, ticketsBought int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticketsBought > 0 && ticketsBought <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// Generates a tickets and sends it to a user e-mail
func sendTicket(ticketsBought int, firstName string, lastname string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v this many tiickets for %v%v", ticketsBought, firstName, lastname)
	println("#########")
	fmt.Printf("Sendingticket: \n %v \n to email addres %v\n", ticket, email)
	println("########")
	wg.Done()
}
