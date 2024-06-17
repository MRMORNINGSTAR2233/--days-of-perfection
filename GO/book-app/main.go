package main

import (
	"fmt"
	"strings"
)

func main() {
	var remainingTickets uint = 50
	const conferenceTickets int = 50
	var conName string = "Go Conference"
	bookings := make([]string, 0, conferenceTickets)

	greetUser(conName)

	for remainingTickets > 0 && len(bookings) < 50 {
		var userTickets uint
		var firstName, lastName string
		var email string

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email address:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidMail := strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidMail && isValidTicket {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("The tickets are sold try next time")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidMail {
				fmt.Println("Please enter a valid email address")
			}
			if !isValidTicket {
				fmt.Println("Please enter a valid number of tickets")
			}
		}
	}
}

func greetUser(conName string) {
	fmt.Printf("Welcome to %v booking application\n", conName)
	fmt.Println("Get your tickets here to attend")
}
