package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaible.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {
		var fisrtName string
		var lastName string
		var email string
		var userTickets uint
		// ask use for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&fisrtName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remainings, so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets
		//bookings[0] = fisrtName + " " + lastName => ARRAY
		bookings = append(bookings, fisrtName+" "+lastName) // Slice

		/*fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The fisrt value: %v\n", bookings[0])
		fmt.Printf("Slice type: %v\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))*/

		fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n", fisrtName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}
