package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]       // var args prend nom du prog
	islash := 0              // var 0 pour commencer boucle à 0
	for i, k := range args { // parcourt le titre args[0]
		if k == '/' { // trouve tous les caractères qui correspondent à '/'
			islash = i //
		}
	}

	for _, k := range args[islash+1:] {
		z01.PrintRune(k)
	}
	z01.PrintRune('\n')
}
