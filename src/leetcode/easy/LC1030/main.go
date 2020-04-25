package main

import (
	"fmt"
	"sort"
)

func main() {
	//[[0,0],[0,1]]
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	//[[0,1],[0,0],[1,1],[1,0]]
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	//[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	index := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res[index] = []int{i, j}
			index++
		}
	}
	sort.Slice(res, func(i, j int) bool {
		a1 := res[i][0] - r0
		a2 := res[i][1] - c0
		a := abs(a1) + abs(a2)
		b1 := res[j][0] - r0
		b2 := res[j][1] - c0
		b := abs(b1) + abs(b2)
		return a < b
	})
	return res
}
func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
