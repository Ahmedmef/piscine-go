package piscine

func Sqrt(nb int) int {
	for x := 0; x <= nb; x++ {
		if x*x == nb {
			return x
		}
	}
	return 0
}
