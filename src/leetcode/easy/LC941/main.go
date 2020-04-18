package main

import "fmt"

func main() {
	fmt.Println(!validMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(!validMountainArray([]int{2, 1}))
	fmt.Println(!validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 1
	for ; i < len(A); i++ {
		if A[i-1] == A[i] {
			return false
		} else if A[i-1] > A[i] {
			break
		}
	}
	if i == 1 || i == len(A) {
		return false
	}
	j := len(A) - 2
	for ; j >= 0; j-- {
		if A[j+1] == A[j] {
			return false
		} else if A[j+1] > A[j] {
			break
		}
	}
	return j+2 == i
}
