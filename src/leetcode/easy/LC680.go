package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("ebcbbececabbacecbbcbe"))
	fmt.Println(!validPalindrome("eeccccbebaeeabebccceea"))
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
}
func validPalindrome(s string) bool {
	length := len(s)
	i := 0
	for ; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return validPalindrome2(s[i+1:length-i]) || validPalindrome2(s[i:length-i-1])
		}
	}
	return true
}
func validPalindrome2(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}
