package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 {
		if remainingTickets == 0 {
			fmt.Println("Our conference is fully booked. come next year")
			break
		}

		lastName, firstName, email, userTickets := helper.GetUserInput()
		isNameValid, isEmailValid, isUserTicketsValid := helper.IsValidateUserInput(lastName, firstName, email, userTickets, remainingTickets)
		if isNameValid && isEmailValid && isUserTicketsValid {
			bookTickets(lastName, firstName, email, userTickets)

			firstNames := getFirstNames()

			fmt.Printf("This people booked tickets: %v\n", firstNames)
		} else {
			if !isEmailValid {
				fmt.Println("Entered email invalid.")
			}
			if !isNameValid {
				fmt.Println("Entered name is too short.")
			}
			if !isUserTicketsValid {
				fmt.Println("Entered tickets number should be more then 0 and less then remaining", remainingTickets)
			}

			fmt.Println("Entered data invalid try again.")
		}
	}

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are avaliable. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, fullName := range bookings {
		var names = strings.Fields(fullName)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func bookTickets(lastName string, firstName string, email string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will get confirmation on your email: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets avaliable. \n", remainingTickets)
}
