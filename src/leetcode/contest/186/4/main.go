package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2) == 23)
	fmt.Println(constrainedSubsetSum([]int{-8269, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790}, 1) == 16091)
	fmt.Println(constrainedSubsetSum([]int{-1, -2, -3}, 1) == -1)
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2) == 37)
}
func constrainedSubsetSum1(nums []int, k int) int {
	max := math.MinInt32
	for j := 0; j < len(nums); j++ {
		m := nums[j]
		for l := 0; l < k && j-l-1 >= 0; l++ {
			v := nums[j] + nums[j-l-1]
			if m < v {
				m = v
			}
		}
		if max < m {
			max = m
		}
		nums[j] = m
	}
	return max
}
func constrainedSubsetSum(nums []int, k int) int {
	res := nums[0]
	var q []int
	for i := 0; i < len(nums); i++ {
		if len(q) != 0 {
			nums[i] += q[0]
		}
		if res < nums[i] {
			res = nums[i]
		}
		for len(q) != 0 && nums[i] > q[len(q)-1] {
			q = q[:len(q)-1]
		}
		if nums[i] > 0 {
			q = append(q, nums[i])
		}
		if i >= k && len(q) != 0 && q[0] == nums[i-k] {
			q = q[1:]
		}
	}
	return res
}
