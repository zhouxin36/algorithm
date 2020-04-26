package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(confusingNumber(6))
	fmt.Println(confusingNumber(89))
	fmt.Println(!confusingNumber(11))
	fmt.Println(!confusingNumber(25))
}
func confusingNumber(num int) bool {
	s := strconv.Itoa(num)
	if strings.Contains(s, "2") || strings.Contains(s, "3") || strings.Contains(s, "4") || strings.Contains(s, "5") || strings.Contains(s, "7") {
		return false
	} else if strings.Contains(s, "6") || strings.Contains(s, "9") {
		return true
	}
	return false
}
