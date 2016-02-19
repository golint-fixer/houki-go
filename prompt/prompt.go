package prompt

import (
	"fmt"
)

func Confirm(prompt string, args ...interface{}) bool {
	s := getStringFromStdin(prompt, args...)
	switch s {
	case "yes", "y", "Y":
		return true
	case "no", "n", "N":
		return false
	default:
		return Confirm(prompt, args...)
	}
}

func getStringFromStdin(prompt string, args ...interface{}) string {
	var s string
	fmt.Printf(prompt, args...)
	fmt.Scanln(&s)
	return s
}
