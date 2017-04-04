package trackTime

import (
	"log"
	"time"
)

// TrackTime takes in start time and logs elapsed run time for function
// Usage: Import trackTime package and call TrackTime() with defer keyword
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("Elapsed time: %s", elapsed)
}
