package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubarraySumCircular([]int{5, 6, 1, 4, 8, -8, 7, -5, 3}) == 29)
	fmt.Println(maxSubarraySumCircular([]int{-2}) == -2)
	fmt.Println(maxSubarraySumCircular([]int{15, -10, 2, -4, 1}) == 16)
	fmt.Println(maxSubarraySumCircular([]int{5, -3, 5}) == 10)
	fmt.Println(maxSubarraySumCircular([]int{-2, -3, -1}) == -1)
	fmt.Println(maxSubarraySumCircular([]int{3, -1, 2, -1}) == 4)
	fmt.Println(maxSubarraySumCircular([]int{1, -2, 3, -2}) == 3)
	fmt.Println(maxSubarraySumCircular([]int{3, -2, 2, -3}) == 3)
}
func maxSubarraySumCircular(A []int) int {
	var sum, qi int
	length := len(A)<<1 - 1
	maxSum := A[0]
	sum = A[0]
	qi = 0
	var f bool
	for i := 1; i < len(A); i++ {
		if A[i] < 0 {
			f = true
		}
		if sum <= 0 {
			sum = A[i]
			qi = i
		} else {
			sum += A[i]
			for qi < i && A[qi] < 0 {
				sum -= A[qi]
				qi++
			}
		}
		if sum > maxSum && qi <= i {
			maxSum = sum
		}
	}
	if !f {
		return maxSum
	}
	for i := len(A); i < length; i++ {
		index := i - len(A)
		if sum <= 0 {
			sum = A[index]
			qi = i
		} else {
			sum += A[index]
			var flag bool
			for (qi < i && A[qi%len(A)] < 0) || i-qi >= len(A) {
				flag = true
				sum -= A[qi%len(A)]
				qi++
			}
			if flag {
				var tsum int
				for j := qi; j < i; j++ {
					tsum += A[j%len(A)]
					if tsum < 0 {
						sum -= tsum
						qi = j + 1
						tsum = 0
					}
				}
			}
		}
		if sum > maxSum && qi <= i {
			maxSum = sum
		}
		if i > len(A) && (sum < 0) {
			break
		}
	}
	return maxSum
}
