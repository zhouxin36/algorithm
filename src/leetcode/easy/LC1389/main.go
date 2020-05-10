package main

import "fmt"

func main() {
	//0,4,1,3,2
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	//0,1,2,3,4
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
	//1
	fmt.Println(createTargetArray([]int{1}, []int{0}))
}
func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if len(res) == index[i] {
			res = append(res, nums[i])
		} else {
			res = append(res, 0)
			copy(res[index[i]+1:], res[index[i]:len(res)-1])
			res[index[i]] = nums[i]
		}
	}
	return res
}
