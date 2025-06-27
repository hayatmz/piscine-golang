package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var st string

	for i := 0; i < len(a); i++ {
		if i >= len(a)-1 {
			st = st + a[i]
		} else {
			st = st + a[i] + "\n"
		}
	}

	for _, b := range st {
		z01.PrintRune(b)
	}

	z01.PrintRune('\n')
}