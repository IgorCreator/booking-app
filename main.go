package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

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

			wg.Add(1)
			go sendTicket(lastName, firstName, email, userTickets)

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

	wg.Wait() // don't give programm fininsh it job until all threads fininsh it job
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are avaliable. \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, userData := range bookings {
		firstNames = append(firstNames, userData.firstName)
	}
	return firstNames
}

func bookTickets(lastName string, firstName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v ticket(s). You will get confirmation on your email: %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("We have %v tickets avaliable. \n", remainingTickets)
}

func sendTicket(lastName string, firstName string, email string, userTickets uint) {

	time.Sleep(10 * time.Second) // simulate delay for confirmation generation

	var ticketConfirm = fmt.Sprintf("%v ticket for %v %v", userTickets, firstName, lastName)

	fmt.Printf("Sending ticket \n\t[%v] \nto email: %v", ticketConfirm, email)

	wg.Done()
}
