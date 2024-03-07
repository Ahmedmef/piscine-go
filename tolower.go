package piscine

func ToLower(s string) string {
	smaller := ""
	for _, i := range s {
		if i >= 'A' && i <= 'Z' {
			smaller += string(i + 32)
		} else {
			smaller += string(i)
		}
	}
	return smaller
}
