package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}) == 3)
	fmt.Println(numSpecialEquivGroups([]string{"abc", "acb", "bac", "bca", "cab", "cba"}) == 3)
}
func numSpecialEquivGroups(A []string) int {
	m := map[[52]int]bool{}
	for _, str := range A {
		count := [52]int{}
		for index, s := range str {
			count[int(s)-'a'+26*(index%2)]++
		}
		m[count] = true
	}
	return len(m)
}
