package main

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}) == 1)
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}) == 2)
}
func minCostToMoveChips(chips []int) int {
	var count int
	for _, v := range chips {
		if v&1 == 1 {
			count++
		}
	}
	if count > len(chips)-count {
		return len(chips) - count
	} else {
		return count
	}
}
