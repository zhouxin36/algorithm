package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}) == 1)
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}) == 1)
}
func peakIndexInMountainArray(A []int) int {
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			return i - 1
		}
	}
	return -1
}
