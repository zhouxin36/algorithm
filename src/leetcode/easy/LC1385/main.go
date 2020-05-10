package main

import "fmt"

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2) == 2)
	fmt.Println(findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3) == 2)
	fmt.Println(findTheDistanceValue([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6) == 1)
}
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	for _, v := range arr1 {
		var flag bool
		for _, vi := range arr2 {
			if abs(v-vi) <= d {
				flag = true
				break
			}
		}
		if !flag {
			res++
		}
	}
	return res
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
