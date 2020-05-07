package main

import "fmt"

func main() {
	//2,4,4,4
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	//1,3,3
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
