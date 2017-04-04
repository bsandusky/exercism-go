package pangram

import (
	"strings"
	"time"

	"../trackTime"
)

const testVersion = 1

// IsPangram takes in a string and outputs a boolean based on whether or not the
// string is a Pangram
func IsPangram(input string) bool {
	defer trackTime.TrackTime(time.Now())

	if len(input) < 26 {
		return false
	}

	input = strings.ToLower(input)

	count := make(map[rune]int)
	for char := 'a'; char < 'z'; char++ {
		count[char] = 0
	}

	for _, char := range input {
		v, p := count[char]
		if p {
			count[char] = v + 1
		}
	}

	for _, v := range count {
		if v == 0 {
			return false
		}
	}
	return true
}
