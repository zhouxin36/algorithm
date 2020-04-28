package main

import "fmt"

func main() {
	fmt.Println(check([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5))
	fmt.Println(!check([]int{10, 100, 101, 101}, 101))
}
func check(nums []int, target int) bool {
	head, tag := 0, len(nums)-1
	for head <= tag {
		mid := (head + tag) >> 1
		if nums[mid] > target {
			tag = mid - 1
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			head = mid
			break
		}
	}
	var count int
	for i := head; i < len(nums) && nums[i] == target; i++ {
		count++
	}
	for i := head - 1; i >= 0 && nums[i] == target; i-- {
		count++
	}
	if len(nums)>>1 < count {
		return true
	} else {
		return false
	}
}
