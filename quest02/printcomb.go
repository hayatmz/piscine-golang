package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for c := '0'; c <= '7'; c++ { // CHIFFRES ET VARIABLES ENTRE ' ' CAR RUNE
		for d := '0'; d <= '8'; d++ {
			for u := '0'; u <= '9'; u++ {
				if c < d && d < u { // CONDITION
					z01.PrintRune(c) // AFFICHE
					z01.PrintRune(d)
					z01.PrintRune(u)
					if c == '7' && d == '8' && u == '9' {
						z01.PrintRune('\n')
						return // ARRETER UNE FONCTION ET RENVOYER DONNEE VERS LA FONCTION QUI L'A APPELE
					}

					z01.PrintRune(',') // RUNE = ' ' & STRING = " "
					z01.PrintRune(' ')
				}
			}
		}
	}
}
