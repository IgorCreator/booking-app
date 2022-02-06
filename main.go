package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are avaliable. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	for remainingTickets > 0 {
		isNoTicketsRemaining := remainingTickets == 0
		if isNoTicketsRemaining {
			fmt.Println("Our conference is fully booked. come next year")
			break
		}

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Print("Enter your name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		fmt.Print("Enter amount of ticket toy want to buy: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Println("You can't booked", userTickets, ". It is more then remaining:", remainingTickets)
		} else {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v ticket(s). You will get confirmation on your email: %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tickets avaliable. \n", remainingTickets)

			firstNames := []string{}
			for _, fullName := range bookings {
				var names = strings.Fields(fullName)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("This people booked tickets: %v\n", firstNames)
		}
	}

	// // Test: arrays
	// var bookingsInintInPlace = [50]string{"Tom", "Jerry", "Spike", "Bob"}
	// fmt.Printf("The whole array: %v\n", bookingsInintInPlace)

	// var bookings [50]string
	// bookings[0] = firstName + " " + lastName
	// bookings[1] = "Jerry"
	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array type %T\n", bookings)
	// fmt.Printf("Array length %v\n", len(bookings))

	// //Test: Slices
	// bookingsArr := []string{}
	// bookingsArr = append(bookingsArr, firstName+" "+lastName)
	// fmt.Printf("The whole slice: %v\n", bookingsArr)
	// fmt.Printf("The first value: %v\n", bookingsArr[0])
	// fmt.Printf("Slice type %T\n", bookingsArr)
	// fmt.Printf("Slice length %v\n", len(bookingsArr))

}
