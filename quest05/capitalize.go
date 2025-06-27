package piscine

func Capitalize(s string) string {
	resultat := ""      // var pour stocker resultat
	xnextlettre := true // variable qui définit si la prochaine lettre doit etre en maj ou min

	for _, x := range s {
		if xnextlettre && x >= 'a' && x <= 'z' {
			// si var xnextletter est comprise entre 'a' et 'z' (table ascii)
			resultat += string(x - 32) // soustraire 32 de la table ascii pour pecho maj ascii
		} else if !xnextlettre && x >= 'A' && x <= 'Z' {
			resultat += string(x + 32) // pareil pour les maj, ajoute 32 pour pecho min dans table ascii
		} else {
			resultat += string(x) // sinon rien changer
		}

		xnextlettre = !((x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') || (x >= '0' && x <= '9'))
		// on vérifie si "x" est en maj, min ou un num et si ce n'est pas le cas la variable "xnextlettre" passe a true ce qui mettra le prochain caractere en capitale et chaque condition est ensuite recheck pour chaque caractere
	}

	return resultat // return nouvelle chaine de caracteres
}
