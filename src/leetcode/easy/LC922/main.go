package main

import (
	"fmt"
)

func main() {
	//"4,5,2,7"
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
}
func sortArrayByParityII(A []int) []int {
	even, odd := 0, 1
	length := len(A)
	for even < length && odd < length {
		for odd < length && A[odd]&1 == 1 {
			odd += 2
		}
		for even < length && A[even]&1 == 0 {
			even += 2
		}
		if odd >= length || even >= length {
			break
		}
		A[odd] ^= A[even]
		A[even] ^= A[odd]
		A[odd] ^= A[even]
	}
	return A
}
