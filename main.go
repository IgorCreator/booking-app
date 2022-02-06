package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("[conferenceName is %T, conferenceTickets is %T, remainingTickets is %T]\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var booking = [50]string{"Tom", "Jerry", "Spike", "Bob"}

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user data
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter amount of ticket toy want to buy: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets
	// userName = "Tom"
	// userTickets = 2
	fmt.Println(firstName, "booked", userTickets, "tickets")
	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will get confirmation on your email: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have total of %v tickets and %v are still avaliable. \n", conferenceTickets, remainingTickets)
}
