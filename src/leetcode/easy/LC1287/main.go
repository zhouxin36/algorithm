package main

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1}) == 1)
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10}) == 6)
}
func findSpecialInteger(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	times := len(arr) >> 2
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			count = 1
		}
		if count > times {
			return arr[i]
		}
	}
	return -1
}
