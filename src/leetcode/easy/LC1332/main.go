package main

import "fmt"

func main() {
	fmt.Println(removePalindromeSub("ababa") == 1)
	fmt.Println(removePalindromeSub("abb") == 2)
	fmt.Println(removePalindromeSub("baabb") == 2)
	fmt.Println(removePalindromeSub("") == 0)
}
func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	bytes := []rune(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	if s == string(bytes) {
		return 1
	} else {
		return 2
	}
}
