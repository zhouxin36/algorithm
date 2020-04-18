package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}) == 3)
	fmt.Println(repeatedNTimes([]int{2, 1, 2, 5, 3, 2}) == 2)
	fmt.Println(repeatedNTimes([]int{5, 1, 5, 2, 5, 3, 5, 4}) == 5)
}
func repeatedNTimes(A []int) int {
	m := map[int]bool{}
	for _, a := range A {
		if m[a] {
			return a
		} else {
			m[a] = true
		}
	}
	return -1
}
