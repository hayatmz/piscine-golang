package piscine

func SplitWhiteSpaces(s string) []string {
	new_count := 0
	size_table := 0
	letter_in_size := 0

	// Count space
	for _, b := range s {
		if string(b) == " " {
			if letter_in_size > 0 {
				size_table++
				letter_in_size = 0
			}
		} else {
			letter_in_size++
		}
	}

	st3 := make([]string, size_table+1)

	letter_in := 0

	for _, b := range s {
		if string(b) == " " {
			if letter_in > 0 {
				new_count++
				letter_in = 0
			}
		} else {
			st3[new_count] = st3[new_count] + string(b)
			letter_in++
		}
	}

	return st3
}