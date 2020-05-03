package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{5, 7, 11, 13}) == 9)
	fmt.Println(missingNumber([]int{15, 13, 12}) == 14)
}
func missingNumber(nums []int) int {
	dif := (nums[len(nums)-1] - nums[0]) / len(nums)
	head, tag := 0, len(nums)-1
	for head < tag {
		mid := (head + tag) >> 1
		if nums[mid] == nums[0]+dif*mid {
			head = mid + 1
		} else {
			tag = mid
		}
	}
	return nums[head] - dif
}
