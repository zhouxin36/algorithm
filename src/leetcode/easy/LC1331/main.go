package main

import (
	"fmt"
	"sort"
)

func main() {
	//4,1,2,3
	fmt.Println(arrayRankTransform([]int{40, 10, 20, 30}))
	//1,1,1
	fmt.Println(arrayRankTransform([]int{100, 100, 100}))
	//5,3,4,2,8,6,7,1,3
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
func arrayRankTransform(arr []int) []int {
	cp := make([]int, len(arr))
	copy(cp, arr)
	sort.Ints(cp)
	m := map[int]int{}
	i := 1
	for _, v := range cp {
		if _, ok := m[v]; !ok {
			m[v] = i
			i++
		}
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = m[arr[i]]
	}
	return arr
}
