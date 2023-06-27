package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 20

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets from here to get attend.\n")

	var userName string
	var userTickets int

	// ask user for their name i.e,, here string and int are the data types

	userName = "Paramesh"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}
