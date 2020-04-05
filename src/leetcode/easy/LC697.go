package main

import (
	"fmt"
)

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}) == 2)
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}) == 6)
}
func findShortestSubArray(nums []int) int {
	stats := map[int]*[3]int{}
	// 0：计数 1：左  2：右
	for i, num := range nums {
		if v, ok := stats[num]; ok {
			v[0]++
			v[2] = i
		} else {
			stats[num] = &[3]int{1, i, i}
		}
	}
	max := 0
	result := len(nums)
	for _, v := range stats {
		length := v[2] + 1 - v[1]
		if v[0] > max {
			result = length
			max = v[0]
		} else if v[0] == max && length < result {
			result = length
		}
	}
	return result
}
