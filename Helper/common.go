// Contains helper methods;
package helper

import "strings"

func ValidateUserInput(firstName string, lastname string, email string, ticketsBought int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastname) >= 2
	isValideEmail := strings.Contains(email, "@")
	isValidTicketNumber := ticketsBought > 0 && ticketsBought <= remainingTickets

	return isValidName, isValidTicketNumber, isValideEmail
}
