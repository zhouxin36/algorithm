package main

import "fmt"

func main() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}) == 2)
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}) == 1)
}
func findNumbers(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if (nums[i] > 9 && nums[i] < 100) || (nums[i] > 999 && nums[i] < 10000) {
			res++
		}
	}
	return res
}
