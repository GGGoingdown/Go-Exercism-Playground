package main

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var message string
	greet := fmt.Sprintf("Welcome to my party, %s!\n", name)
	message += greet

	var digits string
	if table/10 == 0 {
		digits += "0"
	}
	if table/100 == 0 {
		digits += "0"
	}

	sitMsg := fmt.Sprintf("You have been assigned to table %s%d. Your table is %s, exactly %.1f meters from here.\n", digits, table, direction, distance)
	message += sitMsg
	neighborMes := fmt.Sprintf("You will be sitting next to %s.", neighbor)
	message += neighborMes
	return message
}

func main() {
	// msg := AssignTable("Christiane", 7, "Frank", "on the left", 23.7834298)
	// fmt.Println(msg)
	num := 10

	if v := 2 * num; v > 100 {
		fmt.Println("foo")
	} else {
		fmt.Println("Nothing")
	}
	fmt.Println(v)
}
