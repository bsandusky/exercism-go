package accumulate

import (
	"time"

	"../trackTime"
)

const testVersion = 1

// Accumulate takes []string and func(string) string, outputs
//value of evaluation of operation() as []string
func Accumulate(input []string, operation func(string) string) (output []string) {

	defer trackTime.TrackTime(time.Now())

	for _, v := range input {
		output = append(output, operation(v))
	}

	return
}
