package piscine

func IsLower(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}

func IsUpper(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res = true
		} else {
			res = false
			break
		}
	}

	return res
}

func ToUpper(s string) string {
	res := ""
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			res += string(str[i] - 32)
		} else {
			res += string(str[i])
		}
	}
	return res
}

func ToLower(s string) string {
	res := ""
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res += string(str[i] + 32)
		} else {
			res += string(str[i])
		}
	}
	return res
}

func Capitalize(s string) string {
	res := ""
	if IsLower(string(s[0])) {
		res += ToUpper(string(s[0]))
	} else {
		res += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] >= 32 && s[i-1] <= 47 || s[i-1] >= 58 && s[i-1] <= 64 || s[i-1] >= 91 && s[i-1] <= 96 || s[i-1] >= 123 && s[i-1] <= 127 {

			res += ToUpper(string(s[i]))
		} else {
			res += ToLower(string(s[i]))
		}
	}
	return res
}
