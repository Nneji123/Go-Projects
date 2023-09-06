package main

import "fmt"

conferenceName := "Go Conference" // works the same as var but cannot be used to declare variables with types
const conferenceTickets int = 50
var remainingTickets int = 50


func main() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // using %v is for formatting strings
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available") // string concatenation
	fmt.Println("Get your tickets here to attend!") // simple print line

	var userName string
	var userTickets int
	// ask user for their name

	fmt.Scan(userName) // need to add a pointer.
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}


