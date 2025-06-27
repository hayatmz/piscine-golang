package piscine

func AlphaCount(s string) int {
	y := 0

	for _, x := range s {
		if x >= 65 && x <= 90 {
			y++
		}
		if x >= 97 && x <= 122 {
			y++
		}
	}
	return y
}
