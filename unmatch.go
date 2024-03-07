package piscine

func Unmatch(a []int) int {
	for x := 0; x < len(a); x++ {
		for z := 0; z < len(a); z++ {
			if x != z && a[x] == a[z] {
				a[x] = -1
				a[z] = -1
			}
		}
	}
	for x := 0; x < len(a); x++ {
		if a[x] != -1 {
			return a[x]
		}
	}
	return -1
}
