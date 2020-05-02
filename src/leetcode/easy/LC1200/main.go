package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//[[1,2],[2,3],[3,4]]
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
	//[[1,3]]
	fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))
	//[[-14,-10],[19,23],[23,27]]
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	var res [][]int
	dif := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		v := arr[i] - arr[i-1]
		if v == dif {
			res = append(res, []int{arr[i-1], arr[i]})
		} else if v < dif {
			res = res[:0]
			dif = v
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}
