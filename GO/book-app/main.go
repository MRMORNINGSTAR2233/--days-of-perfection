package main

import "fmt"

func main() {
	var remainingTickets uint = 50
	const conferenceTickets int = 10
	var conName string = "Go Conference"
	bookings := [50]string{}

	fmt.Println("Welcome to", conName, "booking application")
	fmt.Println("Get your tickets here to attend")

	var userTickets uint
	var firstName, lastName string
	var email string

	bookings[0] = firstName + " " + lastName

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conName)

	fmt.Printf("These are all our bookings: %v\n", bookings)
}
