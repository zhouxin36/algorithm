package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60) == 58)
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15) == -1)
}
func twoSumLessThanK(arr []int, K int) int {
	sort.Ints(arr)
	start, end := 0, len(arr)-1
	res := -1
	for start < end {
		v := arr[start] + arr[end]
		if v < K {
			if res < v {
				res = v
			}
			start++
		} else {
			end--
		}
	}
	return res
}
