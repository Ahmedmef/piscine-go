package piscine

func DescendAppendRange(max, min int) []int {
	a := []int{}
	for x := max; x > min; x-- {
		a = append(a, x)
	}
	return a
}
