package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}) == 2)
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}) == 10)
}
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	return nums[len(nums)-1]
}
