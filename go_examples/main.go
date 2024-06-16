package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	// Menu
	const StoreName string = "Hezi Ticket Store"
	const TicketCount uint = 50
	var TicketRemain int = 30
	var Bookings[]string
	var FirstName string
	for {

		fmt.Printf("Welcome to our Tickets app \nBuy your tickets here - %s.", StoreName)
		fmt.Printf("\n%v Tickets available from total of %v\n", TicketRemain,TicketCount)

		

		// Request Name and number of tickets
		var UserName string
		var UserTickets int

		// to allow users to enter first and last name together with space we need to use io package
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Please input your full name:\n")
		UserName, err := reader.ReadString('\n')
		UserName = strings.TrimSpace(UserName) // Remove the newline character
		if err != nil {
			fmt.Println("An error occurred when parsing user input:", err)
		}
		// validate user name
		if len(UserName) < 4 {
			fmt.Printf("Invalid full name was provided. Please Register again.\n")	
			continue			
		}
		

		// number of tickets
		fmt.Printf("Please input the number of ticket you want to purchase:\n")
		fmt.Scan(&UserTickets)
		// exit if number of ticket is negative
		if UserTickets < 1 {
			fmt.Printf("Invalid Value of tickets was provided. Please Register again.\n")	
			continue
		}
		// exit if number of ticket is greater than the overall amount of available tickets
		if UserTickets > TicketRemain {
			fmt.Printf("There are not enough tickets available. Current number of available tickets is: %v. Please Register again.\n",TicketRemain)	
			continue
		}

		fmt.Printf("Order of %v Tickets under the name %v successfully processed\n", UserTickets,UserName)

		// Calculate Remain Tickets
		TicketRemain = TicketRemain - UserTickets
		// add new booking
		Bookings = append(Bookings,UserName)

		fmt.Printf("Remaining Tickets Count: %v\n", TicketRemain)
		for index,element := range Bookings {
			if strings.Contains(element," ") { // in case user input first and last name - print only first name
				FirstName = strings.Fields(element)[0]
			} else {
				FirstName = element
			}
			//fmt.Printf("Booking ID: %v Booked for %v\n", index,FirstName)
			fmt.Printf("Booking ID: %v Booked for: %v\n", index,FirstName)
		}
		fmt.Println("================================================================================================")

		// break the loop if we're out of tickets
		if  TicketRemain == 0 {
			fmt.Println("All tickets are sold out. Store is closed")
			break
		}
	}
}
