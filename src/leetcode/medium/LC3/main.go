package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
}
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var arr [128]bool
	var i, j int
	var res int
	for i < len(s) && j < len(s) {
		if arr[s[i]] {
			arr[s[j]] = false
			j++
		} else {
			arr[s[i]] = true
			if i-j > res {
				res = i - j
			}
			i++
		}
	}
	return res + 1
}
