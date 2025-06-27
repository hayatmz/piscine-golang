package piscine

func ToUpper(s string) string {
	x := []byte(s)

	for y := range x {
		if x[y] >= 97 && x[y] <= 122 { // x[y] me permet d'accéder à valeur x d'indice y
			x[y] = x[y] - 32 // MIN 97-122 TO MAJ 65-90
		}
		y++
	}

	return string(x)
}
