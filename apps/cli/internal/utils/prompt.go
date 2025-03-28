package utils

import (
	"fmt"
	"strings"
)

// ConfirmPrompt displays a yes/no prompt with the given message.
// Pressing Enter defaults to "yes".
// Returns true for "yes" and false for "no".
func ConfirmPrompt(message string) bool {
	fmt.Printf("%s [Y/n]: ", message)

	var response string
	fmt.Scanln(&response)

	// Convert to lowercase and trim spaces
	response = strings.TrimSpace(strings.ToLower(response))

	// If empty (user pressed Enter) or starts with 'y', return true
	return response == "" || strings.HasPrefix(response, "y")
}
