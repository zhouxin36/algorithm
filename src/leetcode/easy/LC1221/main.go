package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL") == 4)
	fmt.Println(balancedStringSplit("RLLLLRRRLR") == 3)
	fmt.Println(balancedStringSplit("LLLLRRRR") == 1)
	fmt.Println(balancedStringSplit("RLRRRLLRLL") == 2)
}
func balancedStringSplit(s string) int {
	var stack []int32
	var count int
	for _, v := range s {
		if len(stack) != 0 {
			if stack[0] == v {
				stack = append(stack, v)
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					count++
				}
			}
		} else {
			stack = append(stack, v)
		}
	}
	return count
}
