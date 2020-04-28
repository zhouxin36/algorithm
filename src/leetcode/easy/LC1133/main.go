package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1}) == 8)
	fmt.Println(largestUniqueNumber([]int{9, 9, 8, 8}) == -1)
}
func largestUniqueNumber(nums []int) int {
	res := -1
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, v := range nums {
		if i == 0 && nums[i+1] != v {
			res = v
			break
		} else if i == len(nums)-1 && nums[i-1] != v {
			res = v
			break
		} else if i != 0 && i != len(nums)-1 && nums[i+1] != v && nums[i-1] != v {
			res = v
			break
		}
	}
	return res
}
