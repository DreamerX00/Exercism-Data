package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customer = strings.ToUpper(customer)
    return ("Welcome to the Tech Palace, "+customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    return (strings.Repeat("*",numStarsPerLine)+"\n"+welcomeMsg+"\n"+strings.Repeat("*",numStarsPerLine))
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var newstr string = strings.Trim(oldMsg,"* \n\t")
    return newstr
}
