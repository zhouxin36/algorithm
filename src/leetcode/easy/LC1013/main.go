package main

import (
	"fmt"
)

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
	fmt.Println(!canThreePartsEqualSum([]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}))
	fmt.Println(canThreePartsEqualSum([]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}))
}
func canThreePartsEqualSum(A []int) bool {
	var total int
	var count int
	for _, a := range A {
		total += a
	}
	if total%3 != 0 {
		return false
	}
	num := total / 3
	total = 0
	for _, a := range A[:len(A)-1] {
		total += a
		if total == num {
			count++
			if count == 2 {
				return true
			}
			total = 0
		}
	}
	return false
}
