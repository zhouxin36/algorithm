package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode") == 0)
	fmt.Println(firstUniqChar("loveleetcode") == 2)
}
func firstUniqChar(s string) int {
	var arr [26]int
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
