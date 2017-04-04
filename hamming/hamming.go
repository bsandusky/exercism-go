package hamming

import "errors"

const testVersion = 5

// Distance takes in two strings and outputs int of hamming distance
func Distance(a, b string) (int, error) {

	if len(a) == 0 || len(b) == 0 {
		return 0, nil
	} else if len(a) != len(b) {
		return -1, errors.New("Strings must be of equal length")
	} else {
		var dist int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				dist++
			}
		}
		return dist, nil
	}
}
