package main

//Tech World with Nana Go Learning

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "Booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "tickets are still available.")
	fmt.Println("Get your tickets here.")

}
