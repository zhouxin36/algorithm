package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(rotatedDigits(10) == 4)
}

/**
invalid 3,4,7
self 0,1,8
valid 2,5,6,9
*/
func rotatedDigits(N int) int {
	count := 0
	for i := 1; i <= N; i++ {
		if isValid(i) {
			count++
		}
	}
	return count
}

func isValid(n int) bool {
	v := strconv.Itoa(n)
	flag := false
	for _, a := range v {
		switch a {
		case '3', '4', '7':
			return false
		case '2', '5', '6', '9':
			flag = true
		}
	}
	return flag
}
