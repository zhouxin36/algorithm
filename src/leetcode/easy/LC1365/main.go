package main

import (
	"fmt"
	"sort"
)

func main() {
	//0,3,0,3,0,5
	fmt.Println(smallerNumbersThanCurrent([]int{1, 2, 1, 2, 1, 7}))
	//0,0,2
	fmt.Println(smallerNumbersThanCurrent([]int{1, 1, 7}))
	//2,1,0,3
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	//4,0,1,1,3
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	//0,0,0,0
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
func smallerNumbersThanCurrent(nums []int) []int {
	length := len(nums)
	index := make([]int, length)
	for i := 0; i < length; i++ {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool {
		return nums[index[i]] > nums[index[j]]
	})
	index2 := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		if i < length-1 && nums[index[i]] == nums[index[i+1]] {
			index2[index[i]] = index2[index[i+1]]
		} else {

			index2[index[i]] = i
		}
	}
	for i, v := range index2 {
		index2[i] = length - 1 - v
	}
	return index2
}
func smallerNumbersThanCurrent2(nums []int) []int {

	result := make([]int, 0, 10)

	for _, val := range nums {
		count := 0
		for _, val2 := range nums {
			if val > val2 {
				count++
			}
		}
		result = append(result, count)
	}
	return result

}
