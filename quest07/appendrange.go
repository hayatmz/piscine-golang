package piscine

func AppendRange(min, max int) []int {
	var res []int

	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}
