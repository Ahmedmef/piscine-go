package piscine

func StringToIntSlice(str string) []int {
	var azer []int
	for _, x := range str {
		azer = append(azer, int(x))
	}
	return azer
}
