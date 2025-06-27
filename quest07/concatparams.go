package piscine

func ConcatParams(args []string) string {
	len := len(args)
	var res string
	for i := 0; i < len; i++ {
		res = res + args[i]
		if i != len-1 {
			res = res + "\n"
		}

	}

	return res
}
