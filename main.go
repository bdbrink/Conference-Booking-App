package main

import (
	"fmt"
	"strings"
	"time"
)

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("First name of the booking are: %v \n", firstNames)

			if remainingTickets == 0 {

				fmt.Println("Tickets are booked out for gocoachella.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first or last name is too short, please enter more than 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("email address did entered not have a @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("the number of tickets entered is invalid.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v the better coachella!\n", conferenceName)
	fmt.Printf("We have %v tickets out of %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Tickets for sale")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInputs() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// input information
	fmt.Println("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you want to order: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of booking is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets! Confirmaton is being sent to %v. \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v. \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("#####################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v", ticket, email)
	fmt.Printf("#####################")
}
