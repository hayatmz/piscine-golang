package piscine

func MakeRange(min, max int) []int {
	var table []int

	if max-min <= 0 {
		return table
	}

	table = make([]int, (max - min))

	count := 0

	for i := min; i < max; i++ {

		table[count] = i

		count++

	}

	return table
}