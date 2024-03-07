package piscine

func AppendRange(min, max int) []int {
	a := []int{}
	if min >= max {
		return nil
	}
	for x := min; x < max; x++ {
		a = append(a, x)
	}
	return a
}
