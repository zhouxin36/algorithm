package main

import "fmt"

func main() {
	//[1,4,2,7,5,3,8,6,9]
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	//[1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}))
	//[1,4,2,5,3,8,6,9,7,10,11]
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}))
	//[1,2,3,4,5,6]
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5, 6}}))
}
func findDiagonalOrder(nums [][]int) []int {
	var res []int
	var flags []int
	for i := 0; ; i++ {
		if i < len(nums) {
			flags = append(flags, i)
		}
		for j := len(flags) - 1; j >= 0; j-- {
			v := i - flags[j]
			if v >= len(nums[flags[j]]) {
				copy(flags[j:], flags[j+1:])
				flags = flags[:len(flags)-1]
			} else {
				res = append(res, nums[flags[j]][v])
			}
		}
		if len(flags) == 0 {
			break
		}
	}
	return res
}
