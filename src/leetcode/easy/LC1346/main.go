package main

import "fmt"

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{7, 1, 14, 11}))
	fmt.Println(!checkIfExist([]int{3, 1, 7, 11}))
}
func checkIfExist(arr []int) bool {
	m := map[int]bool{}
	for _, v := range arr {
		if m[2*v] || (v&1 == 0 && m[v>>1]) {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}
