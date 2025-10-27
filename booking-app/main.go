package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func userInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error ocurred when reading the user input: ", err)
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	const conferenceName string = "Go Conference 2025"
	const totalTickets uint32 = 50
	var remainingTickets uint32 = 50
	var bookings = []string{}

	fmt.Printf("Welcome to the %s booking application!\n", conferenceName)
	fmt.Printf("We have a total of %d tickets and %d tickets available for purchase.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend:")

	for {

		name, err := userInput("Insert your full name: ", reader)
		if err != nil {
			fmt.Println("Error while reading name: ", err)
			return
		}

		email, err := userInput("Insert your email: ", reader)
		if err != nil {
			fmt.Println("Error while reading email: ", err)
			return
		}

		userTickets, err := userInput("Insert how many tickets you would like to buy: ", reader)
		ticketsAmount64, err := strconv.ParseUint(userTickets, 10, 32)
		if err != nil {
			fmt.Println("Error while reading amount of tickets: ", err)
			return
		}

		var ticketsAmount32 uint32 = uint32(ticketsAmount64)

		if ticketsAmount32 > remainingTickets {
			fmt.Printf("We only have %d tickets available.\n", remainingTickets)
			continue
		}

		remainingTickets = remainingTickets - ticketsAmount32

		bookings = append(bookings, name)

		fmt.Printf("User %s has booked %d tickets. Purchased tickets are send to %s. \n", name, ticketsAmount32, email)
		fmt.Println(bookings)
		fmt.Println("The amount of remaing tickets is: ", remainingTickets)

		if remainingTickets == 0 {
			// End the application
			fmt.Println("Our converence is currently sold out. Come back next year!")
			break
		}
	}
}
