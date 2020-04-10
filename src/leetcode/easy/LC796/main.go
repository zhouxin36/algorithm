package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(rotateString("abcde", "abcde"))
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(!rotateString("abcde", "abced"))
}
func rotateString2(A string, B string) bool {
	length := len(A)
	if length != len(B) {
		return false
	}
	if length == 0 {
		return true
	}
	if A == B {
		return true
	}
	for i := 1; i < length; i++ {
		j := 0
		for ; j < length; j++ {
			if A[(i+j)%length] != B[j] {
				break
			}
		}
		if j == length {
			return true
		}
	}
	return false
}

func rotateString(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
