package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
	fmt.Println(!isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz"))
	fmt.Println(!isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}
func isAlienSorted(words []string, order string) bool {
	ic, jc := 0, 0
	for index, v := range words[1:] {
		for ic < len(words[index]) && jc < len(v) {
			if words[index][ic] != v[jc] {
				break
			}
			ic++
			jc++
		}
		if ic == len(words[index]) {
			continue
		}
		if jc == len(v) {
			return false
		}
		a := strings.IndexByte(order, words[index][ic])
		b := strings.IndexByte(order, v[jc])
		if a > b {
			return false
		}
	}
	return true
}
