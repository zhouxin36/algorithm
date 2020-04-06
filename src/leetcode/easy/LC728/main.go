package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
func selfDividingNumbers(left int, right int) []int {
	var nums []int
	for i := left; i <= right; i++ {
		if isDN(i) {
			nums = append(nums, i)
		}
	}
	return nums
}

func isDN(num int) bool {
	if num < 10 {
		return true
	}
	n := num
	d := n % 10
	for n != 0 {
		if d == 0 || num%d != 0 {
			return false
		}
		n /= 10
		d = n % 10
	}
	return true
}
