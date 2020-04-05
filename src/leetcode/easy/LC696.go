package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBinarySubstrings("00110011") == 6)
	fmt.Println(countBinarySubstrings("10101") == 4)
}
func countBinarySubstrings(s string) int {
	arr := [...]int{0, 1}
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			if arr[0] > arr[1] {
				count += arr[1]
			} else {
				count += arr[0]
			}
			arr[0], arr[1] = arr[1], 1
		} else {
			arr[1]++
		}
	}
	if arr[0] > arr[1] {
		count += arr[1]
	} else {
		count += arr[0]
	}
	return count
}
