package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}) == 5)
	fmt.Println(largestPerimeter([]int{1, 2, 1}) == 0)
	fmt.Println(largestPerimeter([]int{3, 2, 3, 4}) == 10)
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}) == 8)
}
func largestPerimeter(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] >= A[j]
	})
	var resA []int
	resA = append(resA, A[0], A[1], A[2])
	if resA[2]+resA[1] > resA[0] {
		return resA[0] + resA[1] + resA[2]
	}
	for i := 3; i < len(A); i++ {
		resA = resA[1:]
		resA = append(resA, A[i])
		if resA[2]+resA[1] > resA[0] {
			return resA[0] + resA[1] + resA[2]
		}
	}
	return 0
}
