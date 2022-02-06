package main

import (
	"fmt"
	"strings"
)

func getUserInput() (string, string, string, uint) {
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

	return lastName, firstName, email, userTickets
}

func isValidateUserInput(lastName string, firstName string, email string, userTickets uint) (bool, bool, bool) {
	isNameValid := len(lastName) >= 2 && len(firstName) >= 2
	isEmailValid := strings.Contains(email, "@")
	isUserTicketsValid := userTickets > 0 && userTickets <= remainingTickets

	return isNameValid, isEmailValid, isUserTicketsValid
}

func testsWithArraysAndSlices(firstName string, lastName string) {
	var bookingsInintInPlace = [50]string{"Tom", "Jerry", "Spike", "Bob"}
	fmt.Printf("The whole array: %v\n", bookingsInintInPlace)

	//test arrays
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
