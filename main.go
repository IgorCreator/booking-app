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

	fmt.Println(firstName, "booked", userTickets, "tickets")
	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will get confirmation on your email: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have total of %v tickets and %v are still avaliable. \n", conferenceTickets, remainingTickets)

	// Test: arrays
	var bookingsInintInPlace = [50]string{"Tom", "Jerry", "Spike", "Bob"}
	fmt.Printf("The whole array: %v\n", bookingsInintInPlace)

	var bookings [50]string
	bookings[0] = firstName + " " + lastName
	bookings[1] = "Jerry"
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings))

	//Test: Slices
	bookingsArr := []string{}
	bookingsArr = append(bookingsArr, firstName+" "+lastName)
	fmt.Printf("The whole slice: %v\n", bookingsArr)
	fmt.Printf("The first value: %v\n", bookingsArr[0])
	fmt.Printf("Slice type %T\n", bookingsArr)
	fmt.Printf("Slice length %v\n", len(bookingsArr))

}
