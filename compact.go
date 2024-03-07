package piscine

func Compact(ptr *[]string) int {
	azer := []string{}
	for _, d := range *ptr {
		if d != "" {
			azer = append(azer, d)
		}
	}
	*ptr = azer
	return len(azer)
}
