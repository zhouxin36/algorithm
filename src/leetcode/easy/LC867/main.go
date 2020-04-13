package main

import "fmt"

func main() {
	//[[1,4,7],[2,5,8],[3,6,9]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	//[[1,4],[2,5],[3,6]]
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}
func transpose(A [][]int) [][]int {
	var result [][]int
	for index := range A[0] {
		var t1 []int
		for _, ar := range A {
			t1 = append(t1, ar[index])
		}
		result = append(result, t1)
	}
	return result
}
