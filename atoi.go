package main

import (
	"fmt"
)

func Atoi(s string) int {

	res := 0
	sign := 1
	if len(s) <= 0 {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, x := range s {
		if x < '0' || x > '9' {
			return 0
		}
		res = res*10 + int(x-'0')
	}
	return res * sign
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
