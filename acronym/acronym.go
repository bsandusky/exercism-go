package acronym

import (
	"regexp"
	"strings"
	"time"

	"../trackTime"
)

const testVersion = 2

// Abbreviate takes a string and outputs an acronym as a string
func Abbreviate(input string) (output string) {
	defer trackTime.TrackTime(time.Now())

	regexCamelCase, _ := regexp.Compile(`[a-z]([A-Z])`)
	regexSeparators, _ := regexp.Compile(`[,-;:]`)

	cleaned := regexCamelCase.ReplaceAllString(input, " ${1}")
	cleaned = regexSeparators.ReplaceAllString(cleaned, " ")

	for _, s := range strings.Split(cleaned, " ") {
		if s != "" {
			output += strings.ToUpper(string(s[0]))
		}
	}

	return
}
