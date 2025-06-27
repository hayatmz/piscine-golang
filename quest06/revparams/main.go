package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := len(args) - 1; i > 0; i-- {
		strng := os.Args[i]
		for _, i := range strng {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}
