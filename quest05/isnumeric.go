package piscine

func IsNumeric(s string) bool {
	x := []byte(s)
	for _, y := range x {
		if y >= 48 && y <= 57 {
		} else {
			return false
		}
	}

	return true
}
