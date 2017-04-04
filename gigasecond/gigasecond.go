package gigasecond

import "time"

const testVersion = 4

// AddGigasecond takes a time.Time adds a gigasecond to that time and
// returns new time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000000000000)
}
