package main

import "fmt"

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}) == 1)
	fmt.Println(dominantIndex([]int{1, 2, 3, 4}) == -1)
}
func dominantIndex(nums []int) int {
	var maxIndex int
	for index, n := range nums[1:] {
		if nums[maxIndex] < n {
			maxIndex = index + 1
		}
	}
	for index, n := range nums {
		if maxIndex == index {
			continue
		}
		if n*2 > nums[maxIndex] {
			return -1
		}
	}
	return maxIndex
}
