package bob // package name must match the package name in bob_test.go

import (
	"strings"
	"time"
	"unicode"

	"../trackTime"
)

const testVersion = 2 // same as targetTestVersion

// Hey takes in a string and returns a string response based on
// string characteristics
func Hey(input string) string {
	defer trackTime.TrackTime(time.Now())

	input = strings.TrimSpace(input)

	switch {
	case input == "":
		return "Fine. Be that way!"
	case any(input, unicode.IsUpper) && !any(input, unicode.IsLower):
		return "Whoa, chill out!"
	case input[len(input)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func any(items string, test func(rune) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}
