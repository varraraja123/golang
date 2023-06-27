package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 20

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stll available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets from here to attend.")
}
