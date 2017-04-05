package twelve

import (
	"time"

	"../trackTime"
)

const testVersion = 1

var verses = []struct {
	ordinal string
	verse   string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song takes no argument and generates a string with the song verses
func Song() string {
	defer trackTime.TrackTime(time.Now())

	var song string
	for i := 1; i <= len(verses); i++ {
		song += Verse(i) + "\n"
	}
	return song
}

// Verse takes an int and returns a string with each individual verse
func Verse(x int) string {

	gift := ""

	for i := x; i > 0; i-- {
		gift += verses[i-1].verse

		if i == 2 {
			gift += ", and "
		} else if i > 2 {
			gift += ", "
		}
	}
	return "On the " + verses[x-1].ordinal + " day of Christmas my true love gave to me, " + gift + "."
}
