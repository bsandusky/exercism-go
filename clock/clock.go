package clock

import "fmt"

const testVersion = 4

// Clock data type: int
type Clock int

//New takes two ints: hours and minutes, calculates total time in
//minutes and returns a Clock
func New(h, m int) Clock {
	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

// Add takes an int, adds that value to the total minutes of the Clock object
//and returns a Clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// String returns a string representation of Clock's hours and minutes
//in 00:00 format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
