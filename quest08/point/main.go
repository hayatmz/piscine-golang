package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Nbrstring(x int, y int) {
	var concat string
	var concat1 string

	for x > 0 {
		c := x % 10
		concat = concat + string(c+'0')
		x = x / 10
	}
	for y > 0 {
		c := y % 10
		concat1 = concat1 + string(c+'0')
		y = y / 10
	}
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := len(concat) - 1; i >= 0; i-- {
		z01.PrintRune(rune(concat[i]))
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for i := len(concat1) - 1; i >= 0; i-- {
		z01.PrintRune(rune(concat1[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	points := &point{}
	setPoint(points)

	Nbrstring(points.x, points.y)
}
