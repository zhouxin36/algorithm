package main

import (
	"fmt"
	"sort"
)

func main() {
	//0,1,9,16,100
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	//4,9,9,49,121
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}
