package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var tab []rune        // déclaration tableau de runes
	var verif bool = true // pour verifier que n contient chiffres de 0 à 9
	for n >= 1 {          // boucle pour extraire chiffres de n et add au tableau dans l'ordre inverse
		var x int = n % 10             // obtenir dernier chiffre de n
		tab = append(tab, rune(x+'0')) // conversion chiffre en rune et add au tableau
		n = n / 10                     // diviser nombre par 10 pour passer au prochain chiffre
	}

	for i := '0'; i <= '9'; i++ { // boucle pour imprimer chiffres dans ordre croissant
		for x := range tab { // parcourir tableau runes
			if tab[x] == rune(i) { // vérifier si rune correspond au chiffre actuel de la boucle externe
				z01.PrintRune(tab[x]) // imprimer la rune (chiffre)
				verif = false         // indiquer que le chiffre a été trouvé dans le tableau
			}
		}
	}
	if verif { // si aucun chiffre de 0 à 9 n'a été trouvé, imprimer '0'
		z01.PrintRune('0')
	}
}

// Utilise tableau de runes pour stocker chiffres extraits du nombre.
// Parcourt tableau et imprime les chiffres de 0 à 9 dans l'ordre.
