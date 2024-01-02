package cmd

func countRune(s string, r rune) int {
	var count int
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
