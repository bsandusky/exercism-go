package slice

const testVersion = 1

// All takes int, string returns []string
func All(n int, s string) []string {

	slices := []string{}

	for n <= len(s) {
		slice, ok := First(n, s)
		if ok {
			slices = append(slices, slice)
		} else {
			return []string{}
		}
		s = s[1:]
	}
	return slices
}

// UnsafeFirst takes int, string, returns string
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First takes int, string, returns string, bool
func First(n int, s string) (string, bool) {
	if len(s) < n || n <= 0 {
		return "", false
	}
	return s[:n], true
}
