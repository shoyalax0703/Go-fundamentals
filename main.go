package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidTicketNumber && isValidEmail {

		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		firstNames := getFirstName()
		fmt.Printf("there are all our bookings %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out, Come back next time.")

		} else {
			if !isValidName {
				fmt.Println("first name or last name is invalid, please try again")
			}
			if !isValidEmail {
				fmt.Println("email address is invalid, please try again")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is invalid, please try again")
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets you want:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n", bookings)
	fmt.Printf("thank you %v %v for booking %v tickets and you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("sending your ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
