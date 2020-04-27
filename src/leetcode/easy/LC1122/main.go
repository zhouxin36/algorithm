package main

import (
	"fmt"
	"sort"
)

func main() {
	//[22,28,8,6,17,44]
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
	//2,2,2,1,4,3,3,9,6,7,19
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
}
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := map[int]int{}
	for i, v := range arr2 {
		m[v] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		a1, ok1 := m[arr1[i]]

		a2, ok2 := m[arr1[j]]
		if !ok1 && !ok2 {
			return arr1[i] < arr1[j]
		}
		if !ok1 {
			return false
		}
		if !ok2 {
			return true
		}
		return a1 < a2
	})
	return arr1
}
