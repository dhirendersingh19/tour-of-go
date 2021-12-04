// add 2 string numbers

package main

import "fmt"

func main() {

	s1 := "3423123213"
	s2 := "5412"
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	l := len(s1)
	s1 = reverseString(s1)
	s2 = reverseString(s2)
	var carry byte = 0
	ans1, ans2 := "", ""
	i := 0
	for i = 0; i < l; i++ {
		unit := (s1[i] - 48) + (s2[i] - 48) + carry
		if unit > 9 {
			carry = unit / 10
			unit = (unit % 10)
		}
		unit += 48
		ans1 += string(unit)
	}
	for j := i; j < len(s2); j++ {
		unit := (s2[j] - 48) + carry
		if unit > 9 {
			carry = unit / 10
			unit = (unit % 10)
		}
		unit += 48
		ans2 += string(unit)
	}
	if carry == 1 {
		ans2 += "1"
	}
	fmt.Println(reverseString(ans1 + ans2))
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
