package main

import "fmt"

func main() {
	//[[6],[5],[1],[2],[3],[4],[7]]
	fmt.Println(shiftGrid([][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23))
	//[[9,1,2],[3,4,5],[6,7,8]]
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	//[[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	//[[1,2,3],[4,5,6],[7,8,9]]
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9))
}
func shiftGrid(grid [][]int, k int) [][]int {
	ll := len(grid)
	rl := len(grid[0])
	v := ll * rl
	if v <= k {
		k %= v
	}
	if v == 1 {
		return grid
	}
	nGrid := make([][]int, ll)
	for i := 0; i < ll; i++ {
		nGrid[i] = make([]int, rl)
	}
	for i := 0; i < ll; i++ {
		for j := 0; j < rl; j++ {
			ny := (k / rl) + i
			if (j + (k % rl)) >= rl {
				ny++
			}
			nGrid[ny%ll][(j+(k%rl))%rl] = grid[i][j]
		}
	}
	return nGrid
}
