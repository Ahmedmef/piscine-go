package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	} else if nb > 1 {
		nb = nb * RecursiveFactorial(nb-1)
	}
	return nb
}
