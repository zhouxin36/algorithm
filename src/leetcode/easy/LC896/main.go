package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 2, 3}))
	fmt.Println(isMonotonic([]int{6, 5, 4, 4}))
	fmt.Println(!isMonotonic([]int{1, 3, 2}))
	fmt.Println(isMonotonic([]int{1, 2, 4, 5}))
	fmt.Println(isMonotonic([]int{1, 1, 1}))
}
func isMonotonic1(A []int) bool {
	var inc bool
	index := 1
	for ; index < len(A); index++ {
		if A[index-1] > A[index] {
			inc = false
			break
		} else if A[index-1] < A[index] {
			inc = true
			break
		}
	}
	for ; index < len(A); index++ {
		if inc && A[index] < A[index-1] {
			return false
		}
		if !inc && A[index] > A[index-1] {
			return false
		}
	}
	return true
}
func isMonotonic(A []int) bool {
	inc, dec := true, true
	for index, v := range A[:len(A)-1] {
		if v < A[index+1] {
			dec = false
		} else if v > A[index+1] {
			inc = false
		}
	}
	return inc || dec
}
