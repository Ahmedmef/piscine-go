package piscine

func ReverseMenuIndex(menu []string) []string {
	var azer []string
	for x := len(menu) - 1; x >= 0; x-- {
		azer = append(azer, menu[x])
	}
	return azer
}
