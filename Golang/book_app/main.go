package main

//Tech World with Nana Go Learning

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v Booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here.")

	for remainingTickets > 0 && len(bookings) < 50 {
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

		if customerTickets < int(remainingTickets) {
			remainingTickets = remainingTickets - uint(customerTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you, %v %v for booking %v tickets. The confirmation will be sent to %v\n", firstName, lastName, customerTickets, email)
			fmt.Printf("The remaining tickets in the system is: %v\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are the first names of all of our bookings: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference has reached maxed capacity. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining. Please book within the limits.\n", remainingTickets)
		}

	}

}
