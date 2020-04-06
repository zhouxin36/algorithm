package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}) == 0)
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0}) == 2)
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}) == 3)
	fmt.Println(pivotIndex([]int{1, 2, 3}) == -1)
}
func pivotIndex(nums []int) int {
	sum := 0
	leftSum := 0
	for _, n := range nums {
		sum += n
	}
	for i := 0; i < len(nums); i++ {
		if sum-leftSum-nums[i] == leftSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
