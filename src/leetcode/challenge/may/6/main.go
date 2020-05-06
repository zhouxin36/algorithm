package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}) == 3)
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) == 2)
}
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)>>1]
}
func majorityElement(nums []int) int {
	var res, num int
	for _, v := range nums {
		if num == 0 {
			res, num = v, 1
		} else {
			if res == v {
				num++
			} else {
				num--
			}
		}
	}
	return res
}
