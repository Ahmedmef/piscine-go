package piscine

func ToUpper(s string) string {
	larger := ""
	for _, i := range s {
		if i < 'a' || i > 'z' {
			larger += string(i)
		} else {
			larger += string(i - 32)
		}
	}
	return larger
}
