package piscine

func StrRev(s string) string {
	var Rev string
	for i := len(s) - 1; i >= 0; i-- {
		Rev = Rev + string(s[i])
	}
	return Rev
}
