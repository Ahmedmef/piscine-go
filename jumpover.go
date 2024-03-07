package piscine

func JumpOver(str string) string {
	resul := ""
	if len(str) >= 3 {
		for i, z := range str {
			if i%3 == 2 {
				resul += string(z)
			}
		}
	}
	return resul + "\n"
}
