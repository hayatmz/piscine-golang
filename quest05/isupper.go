package piscine

func IsUpper(s string) bool {
	x := []byte(s)
	for _, valeur := range x {
		if valeur > 90 || valeur < 65 {
			return false
		}
	}

	return true
}
