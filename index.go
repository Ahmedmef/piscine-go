package piscine

func Index(s string, toFind string) int {
	tf := true
	for i := 0; i < len(s); i++ {
		if s[i] == toFind[0] {
			for j := 0; j < len(toFind); j++ {
				if s[i+j] == toFind[j] {
					tf = false
				} else {
					tf = true
					break
				}
			}
			if tf == false {
				return i
			}
		}
	}
	return -1
}
