package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5))      //4
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))             //2
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0)) //3
}
func longestSubarray(nums []int, limit int) int {
	if len(nums) == 1 {
		return 1
	}
	var maxd, mind []int
	var i, j int
	for ; j < len(nums); j++ {
		for len(maxd) != 0 && nums[j] > maxd[len(maxd)-1] {
			maxd = maxd[:len(maxd)-1]
		}
		for len(mind) != 0 && nums[j] < mind[len(mind)-1] {
			mind = mind[:len(mind)-1]
		}
		maxd = append(maxd, nums[j])
		mind = append(mind, nums[j])
		if maxd[0]-mind[0] > limit {
			if maxd[0] == nums[i] {
				maxd = maxd[1:]
			}
			if mind[0] == nums[i] {
				mind = mind[1:]
			}
			i++
		}
	}
	return j - i
}
