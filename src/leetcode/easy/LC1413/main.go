package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}) == 5)
	fmt.Println(minStartValue([]int{1, 2}) == 1)
	fmt.Println(minStartValue([]int{1, -2, -3}) == 5)
}
func minStartValue(nums []int) int {
	res := 1
	sum := res
	for _, v := range nums {
		sum += v
		if sum < 1 {
			res += 1 - sum
			sum = 1
		}
	}
	return res
}
