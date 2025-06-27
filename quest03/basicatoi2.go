package piscine

func BasicAtoi2(s string) int { // transforme string en int
	n := 0                         // pour stocker les int, tu la set up à 0
	for _, ch := range []byte(s) { // parcourt chaque caractère de la chaîne s en tant que tableau de bytes
		ch -= '0'   // redefini la valeur de ch pour retirer le '0' de la table ascii
		if ch > 9 { // vérifie si le chiffre est sup à 9, soit invalide car on veut recup les
			return 0
		}
		n = n*10 + int(ch) // x10 poue décaler mon n à chaque iteration et traduire mon s numerical en byte
	}

	return n
}
