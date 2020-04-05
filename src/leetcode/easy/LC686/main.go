package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab") == 3)
	fmt.Println(repeatedStringMatch("aa", "a") == 1)
	fmt.Println(repeatedStringMatch("abcd", "abcdabcd") == 2)
}
func repeatedStringMatch(A string, B string) int {
	lB := len(B)
	var sb strings.Builder
	sb.WriteString(A)
	q := 1
	for len(sb.String()) < lB {
		sb.WriteString(A)
		q++
	}
	str := sb.String()
	if strings.Contains(str, B) {
		return q
	} else if strings.Contains(str+A, B) {
		return q + 1
	}
	return -1
}

/**
method2  Rabin-Karp (Rolling Hash)
*/
