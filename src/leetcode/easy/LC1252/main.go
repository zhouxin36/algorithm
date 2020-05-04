package main

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}) == 6)
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}) == 0)
}
func oddCells(n int, m int, indices [][]int) int {
	M := make([]bool, m)
	N := make([]bool, n)
	for _, v := range indices {
		M[v[1]] = !M[v[1]]
		N[v[0]] = !N[v[0]]
	}
	var r, c int
	for _, v := range M {
		if v {
			r++
		}
	}
	for _, v := range N {
		if v {
			c++
		}
	}
	return m*c + n*r - 2*c*r
}
