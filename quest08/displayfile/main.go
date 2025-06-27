package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)-1 <= 0 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args)-1 > 1 {
		fmt.Println("Too many arguments")
		return
	}
	x, err := os.ReadFile(os.Args[1:][0])
	if err == nil {
		fmt.Print(string(x))
	}
}
