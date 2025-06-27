package piscine

func Index(s string, toFind string) int {
	sTaille := len(s)
	ToFindTaille := len(toFind)

	for i := 0; i <= sTaille-ToFindTaille; i++ {
		if toFind == s[i:i+ToFindTaille] {
			return (i)
		}
	}

	return -1
}
