package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const shortForm = "1/02/2006 15:04:05"
	t, err := time.Parse(shortForm, date)
	if err != nil {
		panic("Invalid time format")
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const shortForm = "January 2, 2006 15:04:05"
	t, err := time.Parse(shortForm, date)
	if err != nil {
		panic(err.Error())

	}
	now := time.Now()
	if t.Before(now) {
		return true
	}

	return false
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const shortForm = "Monday, January 02, 2006 15:04:05"
	t, err := time.Parse(shortForm, date)
	if err != nil {
		panic(err.Error())

	}
	hour := t.Hour()
	if hour >= 13 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
// Input : "7/25/2019 13:45:00"
// Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	const shortForm = "1/2/2006 15:04:05"
	t, err := time.Parse(shortForm, date)
	if err != nil {
		panic(err.Error())

	}
	s := "You have an appointment on"
	return s + " " + t.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
}
