package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(sortString("aaaabbbbcccc") == "abccbaabccba")
	fmt.Println(sortString("rat") == "art")
	fmt.Println(sortString("leetcode") == "cdelotee")
	fmt.Println(sortString("ggggggg") == "ggggggg")
	fmt.Println(sortString("spo") == "ops")
}
func sortString(s string) string {
	var res strings.Builder
	letter := [26]int{}
	for _, v := range s {
		letter[v-'a']++
	}
	for res.Len() < len(s) {
		for i, v := range letter {
			if v != 0 {
				res.WriteByte(byte('a' + i))
				letter[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if letter[i] != 0 {
				res.WriteByte(byte('a' + i))
				letter[i]--
			}
		}
	}
	return res.String()
}
