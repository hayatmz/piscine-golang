package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 || index == 2 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2) // recup et additionne les 2 index précédents
	}
}