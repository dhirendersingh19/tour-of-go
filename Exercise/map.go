package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	str := strings.Split(s, " ")
	// fmt.Println(str)
	m := make(map[string]int)
	for _, value := range str {
		elem, ok := m[value]
		if ok == false {
			m[value] = 1
		} else {
			m[value] = elem + 1
		}

	}
	return m
}

func main() {
	wc.Test(WordCount)
}
