package piscine

func NRune(s string, n int) rune {
	t := []rune(s)

	if n <= 0 || len(t) < n {
		return 0
	}

	return t[n-1]
}
