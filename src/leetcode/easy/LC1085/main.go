package main

import "fmt"

func main() {
	fmt.Println(sumOfDigits([]int{99, 77, 33, 66, 55}) == 1)
}
func sumOfDigits(A []int) int {
	min := A[0]
	for i := 1; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
	}
	var num int
	for min != 0 {
		num += min % 10
		min /= 10
	}
	if num&1 == 0 {
		return 1
	} else {
		return 0
	}
}
