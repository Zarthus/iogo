package stringtools

func Wrap(s string, maxLength int) []string {
	if maxLength < 0 {
		panic("maxLength must be nonnegative")
	}
	var lines []string
	for i := 0; i < len(s); i += maxLength {
		j := i + maxLength
		if j > len(s) {
			j = len(s)
		}
		lines = append(lines, s[i:j])
	}
	return lines
}
