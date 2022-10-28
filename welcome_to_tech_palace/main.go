package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace," + " " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var starts string
	for i := 0; i < numStarsPerLine; i++ {
		starts += "*"
	}
	message := starts + "\n" + welcomeMsg + "\n" + starts
	return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Replace(oldMsg, "*", "", -1))
}

func main() {
	// name := "eddie"
	// welcome_message := WelcomeMessage(name)
	board_message := AddBorder("Welcome!", 10)
	fmt.Println(board_message)
	fmt.Sprintf()
}

