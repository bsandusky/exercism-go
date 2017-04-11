package secret

const testVersion = 1

// Handshake takes a uint and returns a []string with handshake sequence
func Handshake(x uint) (seq []string) {
	steps := []string{"wink", "double blink", "close your eyes", "jump"}
	seq = make([]string, 0)

	if x <= 0 {
		return
	}

	for s, step := range steps {
		if 1<<uint(s)&x > 0 {
			seq = append(seq, step)
		}
	}

	if 1<<uint(len(steps))&x > 0 {
		reverse(seq)
	}

	return
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
