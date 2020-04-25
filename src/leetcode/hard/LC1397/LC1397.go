package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findGoodStrings(2, "aa", "da", "b") == 51)
	fmt.Println(findGoodStrings(8, "leetcode", "leetgoes", "leet") == 0)
	fmt.Println(findGoodStrings(2, "gx", "gz", "x") == 2)
}
func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	aps := make([]int, n)
	b1 := []byte(s1)
	b2 := []byte(s2)
	e := []byte(evil)
	el := len(evil)
	var sum int
	for i := 0; i < n && sum >= 0; i++ {
		aps[i] = int(s2[i] - s1[i])
		if i >= el-1 && hasBadStr(b1[i-el+1:i+1], b2[i-el+1:i+1], e, el) {
			if i+el < n {

			}
			aps[i]--
		}
		sum = sum*26 + aps[i]
	}
	if !strings.Contains(s2, evil) {
		sum++
	}
	fmt.Println(sum)
	if sum < 0 {
		return 0
	} else {
		return sum
	}
}
func hasBadStr(s1, s2, e []byte, length int) bool {
	for i := 0; i < length; i++ {
		if e[i] > s1[i] && e[i] < s2[i] {
			return true
		}
		if e[i] < s1[i] || e[i] > s2[i] {
			return false
		}
	}
	return true
}
