package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}) == 1)
}
func lastStoneWeight(stones []int) int {
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	for len(stones) > 1 {
		a, b := stones[0], stones[1]
		if a == b {
			stones = stones[2:]
		} else {
			stones = stones[1:]
			stones[0] = a - b
			insertionSort(sort.IntSlice(stones), 0, len(stones))
		}
	}
	if len(stones) == 0 {
		return 0
	} else {
		return stones[0]
	}
}
func insertionSort(data sort.Interface, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && data.Less(j-1, j); j-- {
			data.Swap(j, j-1)
		}
	}
}
