package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{5, 5, 3, 5, 1, -5, 1, -2}, 4))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	NSum(&nums, target, 0, 4, 4, &res, []int{})
	return res
}

func NSum(nums *[]int, target, index, N, ON int, res *[][]int, remark []int) {
	if N == 2 {
		l, r := index, len(*nums)-1
		for l < r {
			if (*nums)[l]+(*nums)[r] > target {
				r--
			} else if (*nums)[l]+(*nums)[r] < target {
				l++
			} else {
				remark = append(remark, []int{(*nums)[l], (*nums)[r]}...)
				r--
				l++
				*res = append(*res, remark)
				for l < r && (*nums)[l-1] == (*nums)[l] {
					l++
				}
				for l < r && (*nums)[r+1] == (*nums)[r] {
					r--
				}
			}
		}
	} else {
		for i := index; i < len(*nums)-N+1; i++ {
			if i > index && (*nums)[i-1] == (*nums)[i] {
				continue
			}
			remark = append(remark, (*nums)[i])
			NSum(nums, target-(*nums)[i], i+1, N-1, ON, res, remark)
			remark = remark[:len(remark)-1]
		}
	}
}
