package booking

import (
	"fmt"
	"time"
)

const layoutOne = "1/2/2006 15:04:05"
const layoutTwo = "Monday, January 2, 2006 15:04:05"
const layoutThree = "January 2, 2006 15:04:05"

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse(layoutOne, date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, err := time.Parse(layoutThree, date)
	if err != nil {
		fmt.Println(err)
	}
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse(layoutTwo, date)
	if err != nil {
		fmt.Println(err)
	}
	return (t.Hour() >= 12 && t.Hour() <= 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t, err := time.Parse(layoutOne, date)
	if err != nil {
		fmt.Println(err)
	}
	return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
