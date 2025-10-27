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
	const totalTickets int = 50
	var remainingTickets int = 30
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
		tickets, err := strconv.ParseUint(userTickets, 10, 0)
		if err != nil {
			fmt.Println("Error while reading amount of tickets: ", err)
			return
		}

		fmt.Printf("User %s has booked %d tickets. Purchased tickets are send to %s. \n", name, tickets, email)

		bookings = append(bookings, name)
		fmt.Println(bookings)
	}
}
