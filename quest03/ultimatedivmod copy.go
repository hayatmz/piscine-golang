package piscine

func UltimateDivMod(a, b *int) {
	division := *a / *b
	restedediv := *a % *b

	*a = division
	*b = restedediv
}

/*func UltimateDivMod(a, b *int) {
}
This function should divide the dereferenced value of a by the dereferenced value of b.
Store the result of the division in the int which a points to.
Store the remainder of the division in the int which b points to.
*/
