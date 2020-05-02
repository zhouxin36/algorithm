package main

import (
	"fmt"
	"sort"
)

func main() {
	//1,5
	fmt.Println(arraysIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}))
}
func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	var res []int
	hit := map[int]int{}
	for _, v := range arr1 {
		hit[v]++
	}
	for _, v := range arr2 {
		if _, ok := hit[v]; ok {
			hit[v]++
		}
	}
	for _, v := range arr3 {
		if _, ok := hit[v]; ok {
			hit[v]++
		}
	}
	for k, v := range hit {
		if v == 3 {
			res = append(res, k)
		}
	}
	sort.Ints(res)
	return res
}
