package piscine

func StrRev(s string) string {
	t := []rune(s) // créer tableau pour séparer lettres du string

	for x, y := 0, len(t)-1; x < y; x, y = x+1, y-1 {
		// len(var)-1 = dernier index d'un tableau | len(var)-2 pour avant-dernier etc..
		// Puisque x et y dans même boucle ils vont se rapprocher ENSEMBLE à chaque boucle

		t[x], t[y] = t[y], t[x]

		// inverser la position de x et y
	}
	return string(t) // ne pose pas de problème pour return en string car c'en était déjà un de base
}
