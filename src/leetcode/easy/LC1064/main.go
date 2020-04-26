package main

import "fmt"

func main() {
	fmt.Println(fixedPoint([]int{-10, -5, 0, 3, 7}) == 3)
	fmt.Println(fixedPoint([]int{-10, -5, 3, 4, 7, 9}) == -1)
}
func fixedPoint(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == i {
			return i
		}
	}
	return -1
}
