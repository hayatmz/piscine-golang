package piscine

func IterativePower(nb int, power int) int { // nb chiffre de base qui va être multiplié par power nombre de fois
	result := 1
	if power < 0 || power > 20 { // si power inférieur a 0 pas possible multiplier, si supp a 20 result trop gros pour memoire
		return 0
	}

	if power == 0 {
		return 1
	}

	for r := 0; r < power; r++ { // temps que r n'égale pas à power, nb va continuer de se multiplier
		result *= nb
	}
	return result
}