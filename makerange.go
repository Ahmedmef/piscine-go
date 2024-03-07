package piscine

func MakeRange(min, max int) []int {
	MaM := max - min
	if MaM <= 0 {
		return nil
	}
	Result := make([]int, MaM)
	for x := range Result {
		Result[x] = min
		min++
	}
	return Result
}
