package piscine

func IsAlpha(s string) bool {
	x := []byte(s)

	for _, y := range x {
		if y >= 65 && y <= 90 || y >= 97 && y <= 122 || y >= 48 && y <= 57 {
			// if y > 65 == !(y <= 65) !(means contraire)
		} else {
			return false // F dans ma boucle pour F si mot pas completement alphanumerical
		}
	}

	return true // en dehors de la boucle car true uniquement si mot alphanumerical donc qui a echappe a la boucle
}
