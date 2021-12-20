package techpalace

import (
	"strings"
)

const WelcomePrompt = "Welcome to the Tech Palace, "

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return WelcomePrompt + strings.ToTitle(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var repeatedStarts = strings.Repeat("*", numStarsPerLine)
	return repeatedStarts + "\n" + welcomeMsg + "\n" + repeatedStarts
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
