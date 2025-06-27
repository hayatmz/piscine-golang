package piscine

func Atoi(s string) int {
	var result int
	signe := 1
	if s != "" {
		if s[0] == '-' { // si premier ch de la chaine est '-'
			signe = -1 // {signe = -1 et n = -1}
			s = s[1:]
		} else if s[0] == '+' {
			s = s[1:]
		}

		for _, x := range s {
			// range s pour vérifier
			if (x < '0') || (x > '9') { // si s n'est pas un num dans table ascii
				return 0
			} else {
				result = result*10 + int(x-'0') // sinon resultat * 10 + int(x-'0')
			}
		}
		return signe * result
	} else {
		return 0
	}
}
