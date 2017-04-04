package leap

const testVersion = 3

// IsLeapYear takes in an int (year) and returns a boolean
// representing whether or not that year is a leap year
func IsLeapYear(z int) bool {

	if (z%4 == 0 && z%100 != 0) || (z%400 == 0) {
		return true
	}
	return false
}
