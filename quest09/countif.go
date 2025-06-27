package piscine

func CountIf(f func(string) bool, tab []string) int {
	panier := 0
	for _, x := range tab {
		if f(x) {
			panier++
		}
	}

	return panier
}