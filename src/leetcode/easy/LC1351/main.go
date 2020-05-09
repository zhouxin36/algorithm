package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}) == 8)
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}) == 0)
	fmt.Println(countNegatives([][]int{{1, -1}, {-1, -1}}) == 3)
	fmt.Println(countNegatives([][]int{{-1}}) == 1)
}
func countNegatives(grid [][]int) int {
	var res int
	for _, r := range grid {
		for i := len(r) - 1; i >= 0; i-- {
			if r[i] < 0 {
				res++
			} else {
				break
			}
		}
	}
	return res
}
