package main

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0) == 0)
	fmt.Println(smallestRangeI([]int{0, 10}, 2) == 6)
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3) == 0)
}
func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, a := range A[1:] {
		if min > a {
			min = a
		} else if max < a {
			max = a
		}
	}
	if max-min-2*K <= 0 {
		return 0
	} else {
		return max - min - 2*K
	}
}
