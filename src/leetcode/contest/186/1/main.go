package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(maxScore("011101") == 5)
	fmt.Println(maxScore("00111") == 5)
	fmt.Println(maxScore("1111") == 3)
}
func maxScore(s string) int {
	var max int
	for i := 0; i < len(s)-1; i++ {
		zero := strings.Count(s[:i+1], "0")
		one := strings.Count(s[i+1:], "1")
		if zero+one > max {
			max = zero + one
		}
	}
	return max
}
