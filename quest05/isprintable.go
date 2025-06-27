package piscine

func IsPrintable(s string) bool {
	x := []byte(s)
	for _, y := range x {
		if y >= 32 && y <= 127 {
		} else {
			return false
		}
	}

	return true
}
