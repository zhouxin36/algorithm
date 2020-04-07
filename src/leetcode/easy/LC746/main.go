package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}) == 6)
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}) == 15)
}
func minCostClimbingStairs(cost []int) int {
	var pref, after int
	for index, c := range cost {
		if index&1 == 0 {
			if pref <= after {
				pref = c + pref
			} else {
				pref = c + after
			}
		} else {
			if pref <= after {
				after = c + pref
			} else {
				after = c + after
			}
		}
	}
	if pref < after {
		return pref
	} else {
		return after
	}
}
