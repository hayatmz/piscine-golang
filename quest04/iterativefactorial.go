package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 20 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	for r := 1; r <= nb; r++ {
		result = r * result
	}
	return result
}