package main

//Tech World with Nana Go Learning

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v Booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")

	var firstName string
	var lastName string
	var email string
	var customerTickets int
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to book: ")
	fmt.Scan(&customerTickets)

	remainingTickets = remainingTickets - uint(customerTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you, %v %v for booking %v tickets. The confirmation will be sent to %v\n", firstName, lastName, customerTickets, email)
	fmt.Printf("The remaining tickets in the system is: %v\n", remainingTickets)

	fmt.Printf("These are all of our bookings: %v\n", bookings)
}
