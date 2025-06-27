package piscine

func IsLower(s string) bool {
	x := []byte(s)

	for _, valeur := range x {
		if valeur > 122 || valeur < 97 {
			return false
		}
	}

	return true
}

// Write a function that returns true if the string
// passed as the parameter contains only lowercase characters
// otherwise, returns false.
