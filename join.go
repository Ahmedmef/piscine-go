package piscine

func Join(strs []string, sep string) string {
	fcb := strs[0]
	for _, rma := range strs[1:] {
		fcb = fcb + sep + rma
	}
	return fcb
}
