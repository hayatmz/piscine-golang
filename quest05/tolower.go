package piscine

func ToLower(s string) string {
	x := []byte(s)

	for y := range x {
		if x[y] <= 90 && x[y] >= 65 {
			x[y] = x[y] + 32
		}
		y++
	}

	return string(x)
}
