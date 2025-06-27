package piscine

func Swap(a, b *int) {
	x := *a
	y := *b

	*a = y
	*b = x
}
