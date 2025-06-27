package piscine

func TrimAtoi(s string) int {
	signe := 1        // varitable pour le signe
	num := 0          // variage pour le nombre extrait de la chaîne
	findnum := false  // bool si chiffre trouvé
	signeneg := false // bool si signe trouvé

	for _, i := range s { // boucle parcourt string
		if i >= '0' && i <= '9' {
			findnum = true
			num = num*10 + int(i-'0')
		} else if i == '-' && !findnum && !signeneg { // ! = inverse de
			signeneg = true
			signe = -1
		}
	}

	return num * signe
}
