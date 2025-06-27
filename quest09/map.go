package piscine

func Map(f func(int) bool, a []int) []bool {
	var x []bool
	for _, y := range a {
		x = append(x, f(y))
	}

	return x
}