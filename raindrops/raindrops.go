package raindrops

import (
	"fmt"
	"time"

	"../trackTime"
)

const testVersion = 2

// Convert takes in an int and outputs a string (Pling, Plong, Plang)
//depending on prime factorization
func Convert(x int) string {

	defer trackTime.TrackTime(time.Now())

	var output string

	if x%3 == 0 {
		output += "Pling"
	}

	if x%5 == 0 {
		output += "Plang"
	}

	if x%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return fmt.Sprintf("%d", x)
	}

	return output
}
