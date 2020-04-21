package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addToArrayForm([]int{0}, 23))
	//1,2,3,4
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	//1,0,0,0,0,0,0,0,0,0,0
	fmt.Println(addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))
	//4,5,5
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
	//1,0,2,1
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
}
func addToArrayForm(A []int, K int) []int {
	index := len(A) - 1
	for K != 0 && index >= 0 {
		k := K % 10
		A[index] += k
		if A[index] > 9 && index > 0 {
			A[index] %= 10
			A[index-1] += 1
		}
		K /= 10
		index--
	}
	if K != 0 {
		if A[0] > 9 {
			A[0] %= 10
			K++
		}
		sk := strconv.Itoa(K)
		ka := make([]int, len(sk))
		for index, v := range sk {
			ka[index] = int(v - '0')
		}
		A = append(ka, A...)
	}
	for index > 0 {
		if A[index] > 9 {
			A[index] %= 10
			A[index-1]++
		}
		index--
	}
	if A[0] > 9 {
		A[0] %= 10
		A = append([]int{1}, A...)
	}
	return A
}
