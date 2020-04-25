package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}) == 3)
}
func heightChecker(heights []int) int {
	arr := make([]int, len(heights))
	copy(arr, heights)
	sort.Ints(arr)
	var count int
	for i := 0; i < len(heights); i++ {
		if arr[i] != heights[i] {
			count++
		}
	}
	return count
}
