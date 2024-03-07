package piscine

func Any(f func(string) bool, arr []string) bool {
	for _, x := range arr {
		if f(x) == true {
			return true
		}
	}
	return false
}
